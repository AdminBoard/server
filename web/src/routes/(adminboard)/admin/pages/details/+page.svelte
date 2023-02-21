<script lang="ts">
    import { page } from '$app/stores';
    import api from '$lib/api';
    import Button from '$lib/components/button.svelte';
    import Content from '$lib/components/content.svelte';
    import notification from '$lib/components/notification/notification';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';
    import { LayoutNode } from '$lib/layout/layout-node';
    import Layout from '$lib/layout/layout.svelte';
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';

    let loading = true;
    let data = { name: '', layout: '' };
    let path: string | null = '';

    let layout: LayoutNode;

    function preview() {
        layout = new LayoutNode().parse(data.layout);
    }

    function save() {
        api.post('admin/page/update?path=' + path, { layout: data.layout })
            .then((resp) => {
                if (resp.status == 0)
                    notification.show('Success', 'Layout saved');
            })
            .catch((e) => {
                notification.error(e);
            });
    }

    onMount(() => {
        const currPath = get(page).url.searchParams.get('path');
        if (currPath == null) return;
        path = currPath;

        api.get('page?path=' + path)
            .then((resp) => {
                if (resp.status == 0) {
                    if (resp.data.layout == null) resp.data.layout = '';
                    data = resp.data;
                }
            })
            .finally(() => {
                loading = false;
            });
    });
</script>

<Titlebar
    >Page: {data?.name}
    {#if path != ''} <div class="path">({path})</div>{/if}</Titlebar
>

<Content>
    <Textfield
        type="textarea"
        disabled={loading}
        class="h-full"
        bind:value={data.layout}
        resizable="vertical"
    />

    <div>
        <Button on:click={preview}>Preview</Button>
        <Button on:click={save}>Save</Button>
    </div>

    <div class="preview">
        <div>Preview:</div>
        <Layout {layout} renderOnly />
    </div>
</Content>

<style lang="scss">
    @import '../../../../../styles/colors';

    .preview {
        @apply m-2 p-2;
        border: 2px solid $color-border;
    }

    .path {
        @apply text-sm ml-1;
        color: transparentize($color-text, 0.3);
    }
</style>
