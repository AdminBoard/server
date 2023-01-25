<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let action: string | null = null;
	export let actionUrl: string | null = null;
	export let actionFields: string | null = null;
	export let onSuccess: any = null;
	export let disabled = false;

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

<button class={clazz} on:click|self={click} {disabled}><slot /></button>
