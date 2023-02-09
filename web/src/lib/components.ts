import Column from "./components/column.svelte"
import Row from "./components/row.svelte"
import DataTable from "./components/data-table.svelte"
import Titlebar from "./components/titlebar.svelte"
import Button from "./components/button.svelte"
import Textfield from "./components/textfield.svelte"
import DataSearch from "./components/data-search.svelte"
import Content from "./components/content.svelte"
import Data from "./components/data.svelte"
import Label from "./components/label.svelte"
import Span from "./components/span.svelte"

const components = { Titlebar, Content, DataTable, Button, Textfield, DataSearch, Column, Row, Data, Label, Span }

export default {
    get,
    has,
}

function get(name: string): any {
    if (name in components) return components[name as keyof typeof components]
    else return null
}

function has(name: string) {
    return (name in components)
}