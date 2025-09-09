import { Component, Element, h, Listen, Prop, Watch } from "@stencil/core";
import { store, Unsubscribe } from "@stencil/redux";
import { AppState } from "../../store/state";
import { ChooserType, ChooserKey, ChooserMode } from "../chooser/chooser";

/**
 * The Choosable component is useful for showing or hiding information based on the
 * currently selected ChooserType and value. For example, a component defined as:
 *
 *     <pulumi-choosable type="language" value="typescript">
 *         I <3 TypeScript.
 *     </pulumi-choosable>
 *
 * ...would display "I <3 TypeScript" only when the user's currently (or most recently)
 * selected language choice is TypeScript.
 */
@Component({
    tag: "pulumi-choosable",
    styleUrl: "choosable.scss",
    shadow: false,
})
export class Choosable {
    private storeUnsubscribe: Unsubscribe;

    @Element()
    el: HTMLElement;

    // The type of chooser to associate with this component instance (e.g., a language chooser).
    @Prop({ mutable: true })
    type: ChooserType;

    // The value to use for determining whether to show or hide this component instance's
    // slotted content.
    @Prop({ mutable: true })
    value: ChooserKey;

    // Similarly to value, this prop allows for providing multiple comma-delimited values.
    @Prop({ mutable: true })
    values: ChooserKey;

    // Choosables are local by default, allowing users to opt into having free-form bits
    // of content simply honor whatever happens to be set on the global store (accepting
    // those bits of content may not show up in all situations).
    @Prop({ mutable: true })
    mode: ChooserMode;

    @Watch("mode")
    onModeChange(newMode: ChooserMode) {
        if (newMode === "local") {
            if (this.storeUnsubscribe) {
                this.storeUnsubscribe();
            }
        }
    }

    // The currently selected value of the supplied chooser type, as .
    @Prop({ mutable: true })
    selection: ChooserKey;

    disconnectedCallback() {
        if (this.storeUnsubscribe) {
            this.storeUnsubscribe();
        }
    }

    @Listen("rendered", { target: "document" })
    onRendered(_event: CustomEvent) {
        // By default, mode is global, until told otherwise by some parental chooser.
        this.mode = "global";

        if (this.mode === "global") {
            this.storeUnsubscribe = store.mapStateToProps(this, (state: AppState) => {
                const {
                    preferences: { language, k8sLanguage, os, cloud, persona, backend },
                } = state;

                switch (this.type) {
                    case "language":
                        return { selection: language };
                    case "k8s-language":
                        return { selection: k8sLanguage };
                    case "os":
                        return { selection: os };
                    case "cloud":
                        return { selection: cloud };
                    case "persona":
                        return { selection: persona };
                    case "backend":
                        return { selection: backend };
                }
            });
        }
    }

    render() {
        const values = this.values ? this.values.split(",").map(v => v.trim()) : [];
        const isActive = this.selection && (this.selection === this.value || values.includes(this.selection));

        // Pages that get structured data (both Variant A and B)
        const structuredDataPages = [
            // Variant A pages
            'authorization/roleassignment',
            'containerservice/managedcluster', 
            'resources/resourcegroup',
            'dbforpostgresql/server',
            // Variant B pages (also get structured data)
            'managedidentity/userassignedidentity',
            'network/virtualnetwork',
            'network/subnet', 
            'app/managedenvironment'
        ];
        
        const needsStructuredData = typeof window !== 'undefined' && 
            structuredDataPages.some(p => 
                window.location.pathname.includes(`azure-native/api-docs/${p}`)
            );
        
        // Apply structured data wrapper for language code blocks on specific pages
        if (this.type === 'language' && needsStructuredData) {
            const languageMap = {
                'typescript': 'TypeScript',
                'javascript': 'JavaScript',
                'python': 'Python',
                'go': 'Go',
                'csharp': 'C#',
                'java': 'Java',
                'yaml': 'YAML'
            };
            
            const langValue = this.value || this.values;
            const mappedLanguage = languageMap[langValue] || langValue;
            
            // Extract resource metadata from URL
            const getResourceMetadata = () => {
                const pathname = typeof window !== 'undefined' ? window.location.pathname : '';
                const match = pathname.match(/azure-native\/api-docs\/([^\/]+)\/([^\/]+)/);
                
                if (match) {
                    const [_, service, resource] = match;
                    // Format service and resource names for display
                    const formatName = (str: string) => {
                        // Handle special cases
                        const specialCases = {
                            'dbforpostgresql': 'Database for PostgreSQL',
                            'containerservice': 'Container Service',
                            'managedidentity': 'Managed Identity',
                            'resourcegroup': 'Resource Group',
                            'roleassignment': 'Role Assignment',
                            'userassignedidentity': 'User Assigned Identity',
                            'virtualnetwork': 'Virtual Network',
                            'managedenvironment': 'Managed Environment',
                            'managedcluster': 'Managed Cluster'
                        };
                        
                        if (specialCases[resource]) {
                            return specialCases[resource];
                        }
                        if (specialCases[service]) {
                            return specialCases[service];
                        }
                        
                        // Default: capitalize first letter
                        return str.charAt(0).toUpperCase() + str.slice(1);
                    };
                    
                    const resourceName = formatName(resource);
                    const serviceName = formatName(service);
                    
                    return {
                        name: `Azure ${serviceName} ${resourceName} - ${mappedLanguage} Example`,
                        description: `Pulumi ${mappedLanguage} code example for creating and managing Azure ${serviceName} ${resourceName} resources using Infrastructure as Code`,
                        keywords: `Azure, ${serviceName}, ${resourceName}, Pulumi, Infrastructure as Code, IaC, ${mappedLanguage}, Cloud, DevOps`,
                        service,
                        resource
                    };
                }
                // Fallback for unexpected URL patterns
                return {
                    name: `Azure Resource - ${mappedLanguage} Example`,
                    description: `Pulumi ${mappedLanguage} code example for Azure resources using Infrastructure as Code`,
                    keywords: `Azure, Pulumi, Infrastructure as Code, IaC, ${mappedLanguage}`,
                    service: 'azure',
                    resource: 'resource'
                };
            };
            
            const metadata = getResourceMetadata();
            
            return (
                <div class={isActive ? "active" : ""}>
                    <div itemscope itemtype="https://schema.org/SoftwareSourceCode">
                        <meta itemprop="name" content={metadata.name} />
                        <meta itemprop="description" content={metadata.description} />
                        <meta itemprop="programmingLanguage" content={mappedLanguage} />
                        <meta itemprop="codeSampleType" content="code snippet" />
                        <meta itemprop="keywords" content={metadata.keywords} />
                        <meta itemprop="runtimePlatform" content="Microsoft Azure" />
                        <meta itemprop="targetProduct" content="Pulumi Infrastructure as Code Platform" />
                        <meta itemprop="codeRepository" content="https://github.com/pulumi/pulumi-azure-native" />
                        <meta itemprop="license" content="https://github.com/pulumi/pulumi-azure-native/blob/master/LICENSE" />
                        <meta itemprop="author" content="Pulumi" />
                        <div itemprop="text">
                            <slot></slot>
                        </div>
                    </div>
                </div>
            );
        }

        // Default rendering for Control pages and non-code content
        return (
            <div class={isActive ? "active" : ""}>
                <slot></slot>
            </div>
        );
    }
}
