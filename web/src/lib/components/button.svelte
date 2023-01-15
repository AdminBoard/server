<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let action: string | null = null;
	export let params: any = null;

	export { clazz as class };

	let clazz = '';

	function click() {
		dispatch('click', { action: action, params: params });
	}
</script>

<button class={clazz} on:click|self={click}><slot /></button>
