import components from "$lib/components"

export class LayoutNode {
    tag: string | null = null;
    name = '';
    text = '';
    attrs: Record<string, any> = [];
    children: LayoutNode[] | null = null;

    constructor(tag?: string) {
        if (tag != null) this.tag = tag;
    }

    parse(str: string): LayoutNode {
        this.children = [];
		const parser = new DOMParser();
		let xml = parser.parseFromString(str, 'text/xml');
		let layout = xml.children.item(0);

        if (layout != null) {
            this.parseChildren(layout.children);
            if (layout.tagName != null) this.tag = layout.tagName;
        }
        
        return this
    }

    parseChildren(children: HTMLCollection) {
        if (children.length == 0) return;
        
        let nodes: LayoutNode[] = [];
        for (let i = 0; i < children.length; i++) {
            let node = new LayoutNode;
            const item = children.item(i)
            if (item != null) {
                if (components.has(item.tagName)) {
                    node.tag = item.tagName
                    if (item.textContent != null) node.text = item.textContent;
                    for (let j = 0; j < item.attributes.length; j++) {
                        const attr = item.attributes.item(j);
                        if (attr != null) {
                            let value = attr.value;
                            
                            if ((value.startsWith('[') && value.endsWith(']')) || (value.startsWith('{') && value.endsWith('}'))) {
                                try {
                                    value = JSON.parse(value)
                                } catch (e) {
                                }
                            }
                            if (attr.nodeName == 'name') node.name = value;
                            else node.attrs[attr.nodeName] = value;
                        }
                    }
                    node.parseChildren(item.children);
                    if (node.children != null) node.text = '';
                } else {
                    node.text = item.outerHTML;
                }
            }
            nodes.push(node);
        }
        // console.log(nodes);
        
        this.children = nodes;
    }
}

