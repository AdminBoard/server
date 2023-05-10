<script lang="ts">
	import api from '$lib/api';
	import { createEventDispatcher, onMount } from 'svelte';
	import Table from './table.svelte';

	export let src = '';
	export let columns: any[] = [];
	export let selectable = false;
	export let multiple = false;
	export let data: any[] = [];
	export let selections: any[] = [];

	let page = { offset: 0, count: 100 };

	const dispatch = createEventDispatcher();

	export function refresh() {
		if (src != '') {
			if (src.startsWith('/api/')) src = src.substring(4);
			api.post(src, { page: page })
				.then((resp) => {
					if (resp.status == 0 && resp.data != null) data = resp.data;
				})
				.catch((e) => dispatch('error', e));
		}
	}

	onMount(() => refresh());
</script>

<Table
	bind:columns
	bind:data
	bind:selectable
	on:select
	bind:multiple
	bind:selections
/>
<slot />
