import { Component } from "solid-js";
import { marked } from "marked";
import guideMarkdown from "../../../SUBSCRIPTION_INSTRUCTION.md?raw";

const guideHtml = marked.parse(guideMarkdown, { async: false });

export const Guide: Component = () => {
    return (
        <main class="guide">
            <header class="guide-header">
                <a href="/signup">Naarad Signup</a>
            </header>
            <article class="guide-content" innerHTML={guideHtml}></article>
        </main>
    );
};