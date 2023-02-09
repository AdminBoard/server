<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let action: string | null = null;
	export let actionUrl: string | null = null;
	export let actionFields: string | null = null;
	export let onSuccess: any = null;
	export let disabled = false;
	export let tooltip: string | null = null;

	export { clazz as class };

	let clazz = '';

	function click() {
		dispatch('click', {
			action: action,
			actionUrl: actionUrl,
			actionFields: actionFields,
			onSuccess: onSuccess,
		});
	}
</script>

<button class={clazz} data-tooltip={tooltip} on:click={click} {disabled}
	><slot /></button
>

<style lang="scss">
	button {
		& > :global(*) {
			@apply align-middle;
			@apply inline-block;
		}
	}
</style>
