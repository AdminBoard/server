<script lang="ts">
	import { page } from '$app/stores';
	import api from '$lib/api';
	import Titlebar from '$lib/components/titlebar.svelte';
	import { onDestroy, onMount } from 'svelte';
	import { get, type Subscriber, type Unsubscriber } from 'svelte/store';
	import MenuPage from './menu/menu-page.svelte';
	import PagesPage from './pages/pages-page.svelte';

	let path = '';
	let layout = { kind: '' };
	let error = '';

	function refresh(newPath: string) {
		error = '';
	}

	let unsubscribe = page.subscribe((data) => {
		api.get('page?url=/admin/' + data.params.slug).then((resp) => {
			if (resp.status == 0) {
				if (resp.data.layout != null)
					layout = JSON.parse(resp.data.layout);
			} else {
				error = resp.message;
			}
		});
	});

	onDestroy(unsubscribe);
</script>

{#if error != ''}
	<h2>{error}</h2>
{/if}

{#if layout.kind == 'admin.menu'}
	<MenuPage />
{:else if layout.kind == 'admin.pages'}
	<PagesPage />
{/if}
