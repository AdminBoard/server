<script lang="ts">
    import { goto } from '$app/navigation';
    import { SelectAction } from '$lib/class/select-action';
    import Button from '$lib/components/button.svelte';
    import Content from '$lib/components/content.svelte';
    import DataTable from '$lib/components/data-table.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';

    function select(ev: any) {
        goto('/admin/apis/details/' + ev.detail.id);
    }
</script>

<Titlebar
    ><span>APIs</span><Button
        class="material-icons"
        on:click={() => goto('/admin/apis/add')}>add</Button
    ></Titlebar
>
<Content>
    <DataTable
        source="/admin/apis"
        columns={[
            { name: 'url', label: 'Path' },
            { name: 'method', label: 'Method' },
            { name: 'secure', label: 'Secure' },
            { name: 'description', label: 'Description' },
            { name: 'status', label: 'Status' },
        ]}
        selectable={true}
        on:select={select}
    />
</Content>
