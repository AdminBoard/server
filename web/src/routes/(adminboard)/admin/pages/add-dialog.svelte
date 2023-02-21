<script lang="ts">
    import Button from '$lib/components/button.svelte';
    import Content from '$lib/components/content.svelte';
    import Dialog from '$lib/components/dialog.svelte';
    import Row from '$lib/components/row.svelte';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';
    import { createEventDispatcher } from 'svelte';

    export let visible = false,
        disabled = false;

    let path = '',
        name = '',
        description = '';

    let dispath = createEventDispatcher();

    export function show() {
        disabled = false;
        path = '';
        name = '';
        description = '';
        visible = true;
    }
    export function dismiss() {
        path = '';
        name = '';
        description = '';
        visible = false;
    }
</script>

<!-- Add Form -->
<Dialog bind:visible>
    <Titlebar>Add Page</Titlebar>
    <Content>
        <Textfield bind:value={path} label="Path" bind:disabled />
        <Textfield bind:value={name} label="Title" bind:disabled />
        <Textfield
            bind:value={description}
            label="Description"
            type="textarea"
            bind:disabled
        />
        <Row>
            <div class="flex-1" />
            <Button on:click={dismiss} bind:disabled>Cancel</Button>
            <Button
                on:click={() =>
                    dispath('save', {
                        path: path,
                        name: name,
                        description: description,
                    })}
                bind:disabled
            >
                Save
            </Button>
        </Row>
    </Content>
</Dialog>
