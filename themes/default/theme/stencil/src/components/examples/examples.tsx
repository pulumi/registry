import { Component, Element, h } from "@stencil/core";

@Component({
    tag: "pulumi-examples",
    styleUrl: "examples.scss",
    shadow: false,
})
export class Examples {
    @Element()
    el: HTMLElement;

    componentWillRender() {
        const headings = Array.from(this.el.querySelectorAll("pulumi-examples h3"));

        headings.forEach(heading => {
            // Prepend a span to each heading. We use these to hold the caret symbols.
            const span = document.createElement("span");
            heading.prepend(span);

            // Listen for clicks on headings.
            heading.addEventListener("click", event => {
                const clicked = event.target as HTMLElement;

                // Ignore anchor clicks; we want them to be clickable without triggering
                // any expand/collapse behavior.
                if (clicked.classList.contains("anchorjs-link")) {
                    return;
                }

                // When a heading is clicked, "toggle" it.
                this.toggle(heading);
            });
        });

        // Expand the first example by default.
        if (headings && headings.length > 0) {
            this.toggle(headings[0]);
        }
    }

    componentDidLoad() {
        // Only Variant B pages get the crawler-visible canonical
        const variantBPages = [
            'managedidentity/userassignedidentity',
            'network/virtualnetwork',
            'network/subnet', 
            'app/managedenvironment'
        ];
        
        const isVariantB = typeof window !== 'undefined' && 
            variantBPages.some(p => 
                window.location.pathname.includes(`azure-native/api-docs/${p}`)
            );
        
        if (!isVariantB) {
            return;
        }
        
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
                    serviceName,
                    resourceName,
                    service,
                    resource,
                    pageUrl: typeof window !== 'undefined' ? window.location.href : ''
                };
            }
            // Fallback for unexpected URL patterns
            return {
                serviceName: 'Azure',
                resourceName: 'Resource',
                service: 'azure',
                resource: 'resource',
                pageUrl: typeof window !== 'undefined' ? window.location.href : ''
            };
        };
        
        const resourceMetadata = getResourceMetadata();
        
        // Create crawler-visible container
        const crawlerDiv = document.createElement('div');
        crawlerDiv.className = 'crawler-canonical-examples';
        crawlerDiv.style.cssText = 'position: absolute; left: -10000px; width: 1px; height: 1px; overflow: hidden;';
        
        // Language mapping for structured data
        const languageMap = {
            'typescript': 'TypeScript',
            'javascript': 'JavaScript', 
            'python': 'Python',
            'go': 'Go',
            'csharp': 'C#',
            'java': 'Java',
            'yaml': 'YAML'
        };
        
        // Extract all example titles from h3 headings within this component
        const exampleTitles = Array.from(this.el.querySelectorAll('h3')).map(h3 => h3.textContent?.trim() || '');
        
        // Extract all language versions from the DOM
        const languages = ['typescript', 'python', 'go', 'csharp', 'java', 'yaml'];
        
        languages.forEach(lang => {
            const choosables = this.el.querySelectorAll(`pulumi-choosable[values="${lang}"], pulumi-choosable[value="${lang}"]`);
            
            choosables.forEach((choosable, index) => {
                const codeElement = choosable.querySelector('pre code');
                if (codeElement && codeElement.textContent) {
                    const mappedLanguage = languageMap[lang] || lang;
                    const exampleTitle = exampleTitles[Math.floor(index / languages.length)] || 'Example';
                    
                    const wrapper = document.createElement('div');
                    wrapper.setAttribute('itemscope', '');
                    wrapper.setAttribute('itemtype', 'https://schema.org/SoftwareSourceCode');
                    
                    // Add comprehensive metadata
                    const metaElements = [
                        { prop: 'name', content: `Azure ${resourceMetadata.serviceName} ${resourceMetadata.resourceName} - ${exampleTitle} - ${mappedLanguage}` },
                        { prop: 'description', content: `Pulumi ${mappedLanguage} code example for ${exampleTitle} with Azure ${resourceMetadata.serviceName} ${resourceMetadata.resourceName}. This example demonstrates Infrastructure as Code patterns for Azure resources.` },
                        { prop: 'programmingLanguage', content: mappedLanguage },
                        { prop: 'codeSampleType', content: 'code snippet' },
                        { prop: 'keywords', content: `Azure, ${resourceMetadata.serviceName}, ${resourceMetadata.resourceName}, ${exampleTitle}, Pulumi, Infrastructure as Code, IaC, ${mappedLanguage}, Cloud, DevOps` },
                        { prop: 'runtimePlatform', content: 'Microsoft Azure' },
                        { prop: 'targetProduct', content: 'Pulumi Infrastructure as Code Platform' },
                        { prop: 'codeRepository', content: 'https://github.com/pulumi/pulumi-azure-native' },
                        { prop: 'license', content: 'https://github.com/pulumi/pulumi-azure-native/blob/master/LICENSE' },
                        { prop: 'author', content: 'Pulumi' },
                        { prop: 'isPartOf', content: resourceMetadata.pageUrl }
                    ];
                    
                    metaElements.forEach(({ prop, content }) => {
                        const meta = document.createElement('meta');
                        meta.setAttribute('itemprop', prop);
                        meta.setAttribute('content', content);
                        wrapper.appendChild(meta);
                    });
                    
                    // Add the code content
                    const pre = document.createElement('pre');
                    const code = document.createElement('code');
                    code.setAttribute('itemprop', 'text');
                    code.textContent = codeElement.textContent;
                    
                    pre.appendChild(code);
                    wrapper.appendChild(pre);
                    
                    crawlerDiv.appendChild(wrapper);
                }
            });
        });
        
        // Only append if we found code snippets
        if (crawlerDiv.children.length > 0) {
            this.el.appendChild(crawlerDiv);
        }
    }

    // Show or hide the examples associated with each heading.
    private toggle(heading: Element) {
        heading.classList.toggle("expanded");

        // Examples are contained within divs alongside their corresponding headings, so
        // we start with the heading that was clicked, then iterate over the siblings that
        // follow, toggling them also, until we come to one that isn't a div.
        let example = heading.nextElementSibling;

        while (example) {
            if (example.nodeName === "DIV") {
                example.classList.toggle("expanded");
                example = example.nextElementSibling;
            } else {
                break;
            }
        }
    }

    render() {
        return (
            <div>
                <slot></slot>
            </div>
        );
    }
}
