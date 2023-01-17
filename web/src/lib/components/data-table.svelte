<script lang="ts">
	import { goto } from '$app/navigation';
	import api from '$lib/api';
	import { onMount } from 'svelte';
	import { showPage } from './modal.svelte';
	import Table from './table.svelte';

	export let source: string = '';
	export let columns: any[] = [];
	export let onSelect: any = null;

	let page = { offset: 0, count: 100 };

	let rows: any[] = [];

	onMount(() => {
		if (source != '') {
			if (source.startsWith('/api/')) source = source.substring(4);
			api.post(source, { page: page }).then((resp) => {
				if (resp.status == 0) rows = resp.data;
			});
		}
	});

	function select(ev: any) {
		const detail = ev.detail;

		switch (onSelect.action) {
			case 'popup':
				showPage('', onSelect.url, detail);
				break;
			default:
				if (onSelect.url != '') {
					let url = onSelect.url;
					for (const key in detail) {
						url = url.replaceAll('%' + key + '%', detail[key]);
					}
					goto(url);
				}
		}
	}
</script>

<Table {columns} {rows} selectable={onSelect != null} on:select={select} />
<slot />
