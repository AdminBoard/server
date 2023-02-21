<script lang="ts">
    import api from '$lib/api';
    import Button from '$lib/components/button.svelte';
    import Column from '$lib/components/column.svelte';
    import Content from '$lib/components/content.svelte';
    import Dialog from '$lib/components/dialog.svelte';
    import notification from '$lib/components/notification/notification';
    import Row from '$lib/components/row.svelte';
    import Textfield from '$lib/components/textfield.svelte';
    import Titlebar from '$lib/components/titlebar.svelte';
    import { createEventDispatcher } from 'svelte';

    export let visible = false;

    let id = 0,
        parameter = '',
        property = '',
        query = '';

    let disabled = false;
    let dispatch = createEventDispatcher();

    export function show(data: any) {
        id = data.id;
        parameter = data.parameter;
        property = data.property;
        query = data.query;
        visible = true;
    }

    export function dismiss() {
        id = 0;
        parameter = '';
        property = '';
        query = '';
        visible = false;
        disabled = false;
    }

    export function save() {
        disabled = true;
        api.postData('/api/admin/apis/query/update?id=' + id, {
            property: property,
            query: query,
            parameter: parameter,
        })
            .then(() => {
                notification.show('Success', 'Operation success');
                dispatch('success');
                dismiss();
            })
            .catch((e) => notification.error(e))
            .finally(() => (disabled = false));
    }
</script>

<Dialog bind:visible>
    <Titlebar>Query</Titlebar>
    <Content>
        <Row>
            <Textfield bind:value={property} label="Property" />
            <Column class="flex-1">
                <Textfield
                    type="textarea"
                    label="Query"
                    bind:value={query}
                    resizable="both"
                />
                <Textfield label="Parameters" bind:value={parameter} />
            </Column>
        </Row>
        <Row>
            <div class="flex-1" />
            <Button on:click={dismiss}>Cancel</Button>
            <Button on:click={save}>Save</Button>
        </Row>
    </Content>
</Dialog>
