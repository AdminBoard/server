<script lang="ts">
	import api from '$lib/api';
	import { onMount } from 'svelte';
	import Table from './table.svelte';

	export let src = '';
	export let columns: any[] = [];
	export let selectable = false;
	export let multiple = false;

	let page = { offset: 0, count: 100 };

	let data: any[] = [];

	onMount(() => {
		if (src != '') {
			if (src.startsWith('/api/')) src = src.substring(4);
			api.post(src, { page: page }).then((resp) => {
				if (resp.status == 0 && resp.data != null) data = resp.data;
			});
		}
	});
</script>

<Table bind:columns bind:data bind:selectable on:select bind:multiple />
<slot />
