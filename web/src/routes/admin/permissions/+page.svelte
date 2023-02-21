<script lang="ts">
    import { goto } from '$app/navigation';
    import Button from '$lib/components/button.svelte';
    import DataTable from '$lib/components/data-table.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';

    import Layout from '$lib/layout/layout.svelte';
</script>

<Titlebar
    ><div>Permissions</div>
    <Button
        class="material-icons"
        on:click={() => goto('/admin/permissions/details')}>add</Button
    ></Titlebar
>
<DataTable
    src="/api/admin/permissions"
    columns={[
        { name: 'name', label: 'Name' },
        { name: 'description', label: 'Description' },
    ]}
    on:select={(ev) => goto('/admin/permissions/details?id=' + ev.detail.id)}
    selectable
/>
<!-- onSelect={{
    on:select={(ev: any) => goto('/admin/permissions/details?id=' + ev.id)}
        action: 'open',
        actionUrl: '/admin/permissions/details?id=[id]',
    }} -->
