<script lang="ts">
    import api from '$lib/api';
    import { onMount } from 'svelte';
    import Textfield from './textfield.svelte';

    export let label = '';
    export let placeholder = '';
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

    let isFocus = false;

    onMount(() => {
        oldValue = value;
    });

    export function focus() {
        textfield.focus();
    }

    function setFocus(focus: boolean) {
        if (focus) isFocus = true;
        else
            setTimeout(() => {
                isFocus = false;
            }, 200);
    }

    function save() {
        if (saveUrl == '') return;
        disabled = true;

        let params: Record<string, any> = {};
        params[name] = value;

        api.post(saveUrl, params)
            .then((resp) => {
                if (resp.status == 0) {
                    saveCallback!(null);
                    oldValue = value;
                } else saveCallback!(resp.message);
            })
            .finally(() => (disabled = false));
    }
</script>

<div class="input {clazz}" style="width: {width}rem;">
    <Textfield
        bind:this={textfield}
        bind:value
        bind:label
        bind:placeholder
        on:input
        on:keyup
        on:focus={() => setFocus(true)}
        on:blur={() => setFocus(false)}
        class={clazz}
        {disabled}
    />
    <div class="buttons" class:invisible={value == oldValue || !isFocus}>
        <button class="clear" on:click={() => (value = oldValue)} {disabled}
            ><i class="material-icons">undo</i></button
        >
        <button class="done" on:click={save} {disabled}
            ><i class="material-icons">done</i></button
        >
    </div>
</div>

<style lang="scss">
    @import '../../styles/colors';
    .input {
        position: relative;
        & .buttons {
            @apply absolute bottom-0 right-1 h-full transition-all;

            transform: translateX(100%);

            &.invisible {
                opacity: 0;
            }

            & button {
                @apply h-full text-sm rounded-none;
                border-top: 1px solid $color-secondary;
                border-bottom: 1px solid $color-secondary;
                background-color: $color-primary;
                color: $color-text;
                &.done {
                    @apply rounded-r;
                    margin: -0.25rem;
                }
            }
        }
    }
</style>
