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

    let name = '';
    let loading = true;
    let data = { layout: '' };

    let layout: LayoutNode;

    onMount(() => {
        const paths = get(page).url.pathname.split('/');
        name = paths[paths.length - 1];

        api.get('page?name=' + name)
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

    function preview() {
        layout = new LayoutNode().parse(data.layout);
    }

    function save() {
        api.post('admin/page/update?name=' + name, { layout: data.layout })
            .then((resp) => {
                if (resp.status == 0)
                    notification.show('Success', 'Layout saved');
            })
            .catch((e) => {
                notification.show('Error', e);
            });
    }
</script>

<Titlebar>Page: {name}</Titlebar>

<Content>
    <Textfield
        type="textarea"
        disabled={loading}
        class="h-full"
        bind:value={data.layout}
        resizeable="vertical"
    />

    <div>
        <Button on:click={preview}>Preview</Button>
        <Button on:click={save}>Save</Button>
    </div>

    <div class="preview">
        <div>Preview:</div>
        <Layout {layout} />
    </div>
</Content>

<style lang="scss">
    @import '../../../../../styles/colors';

    .preview {
        @apply m-2 p-2;
        border: 2px solid $color-border;
    }
</style>
