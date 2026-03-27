import { defineCustomElements } from "../../stencil/dist/loader";

import "../scss/main.scss";

import "./misc";
import "./nav";
import "./chooser";
import "./noselect";
import "./tracking";
import "./copybutton";
import "./code-tabbed";
import "./packages";
import "./toc";
import "./pulumi-ai-menu";
import "./docs-main";
import "./algolia/autocomplete";

// Register all Stencil components.
defineCustomElements();
