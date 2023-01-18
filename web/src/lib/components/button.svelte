<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let action: string | null = null;
	export let params: any = null;
	export let disabled = false;

	export { clazz as class };

	let clazz = '';

	function click() {
		dispatch('click', { action: action, params: params });
	}
</script>

<button class={clazz} on:click|self={click} {disabled}><slot /></button>
