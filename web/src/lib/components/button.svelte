<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let action: string | null = null;
	export let params: any = null;

	export { clazz as class };

	let clazz = '';

	function click() {
		dispatch('click', { action: action, params: params });
		// switch (action) {
		// 	case 'popup':
		// 		console.log(params);

		// 		// try {
		// 		// 	if (params != null) {
		// 		// 		show(params.title, params.url);
		// 		// 	}
		// 		// } catch (e) {}
		// 		// show()
		// 		break;
		// 	default:
		// 		dispatch('click', { action: action });
		// }
	}
</script>

<button class={clazz} on:click|self={click}><slot /></button>
