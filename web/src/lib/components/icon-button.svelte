<script lang="ts">
    import Button from './button.svelte';

    export let data: any = null;
    export let disabled = false;
    export let tooltip: string | null = null;
    export { clazz as class };

    let clazz = '';
</script>

<Button
    bind:data
    bind:disabled
    bind:tooltip
    on:click
    class={(clazz + ' material-icons text-base').trim()}><slot /></Button
>
