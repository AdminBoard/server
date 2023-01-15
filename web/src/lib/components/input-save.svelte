<script lang="ts">
	import api from "$lib/api";
	import { onMount } from "svelte";
	import Textfield from "./textfield.svelte";

	export let name = '';
	export let value = '';
	export let width = '16';
	export let disabled = false;
	export let allowCancel = false;
	export let saveUrl = '';
	export let saveCallback: Function | null = null;

    export { clazz as class };
    let clazz = '';

    let oldValue = '';
    let textfield: Textfield;

    onMount(() => {
        oldValue = value
    })

    function save() {
        if (saveUrl == '') return;
        disabled = true;

        let params: Record<string, any> = {};
        params[name] = value;

        api.post(saveUrl, params).then((resp) => {
            if (resp.status == 0) {
                saveCallback!(null);
                oldValue = value;
            } else saveCallback!(resp.message);
        }).finally(() => (disabled = false));

    }

</script>
<div class="input {clazz}">
    <Textfield bind:this={textfield} bind:value={value} class={clazz} {disabled} />
    <div class="button" class:invisible={value == oldValue}>
    <button class="material-icons" on:click={() => (value = oldValue)} {disabled}>undo</button>
    <button class="material-icons" on:click={save} {disabled}>done</button>
    </div>
</div>

<style lang="scss">
    @import '../../styles/colors';
    .input {
        position: relative;
        & .button {
            @apply space-x-0 absolute;
            top: 1px;
            right: 1px;
            bottom: 1px;

            & button {
                @apply px-2 h-full;
                background-color: $color-border;
                color: $color-primary;
            }
        }
    }
</style>