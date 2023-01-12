<script lang="ts">
	import '../app.scss';

	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { load, loggedIn } from '$lib/session';
	import Sidebar from '$lib/components/sidebar.svelte';
	import Notification from '$lib/components/notification/notification.svelte';
	import Modal from '$lib/components/modal.svelte';
	import { page } from '$app/stores';
	import { get } from 'svelte/store';

	export const ssr = false;

	let loadSession = false;

	onMount(async () => {
		load()
			.then((resp) => {
				switch (resp.status) {
					case 1:
						if (get(page).url.pathname != '/user/change_password') goto('/user/change_password');
						break;
				}
			})
			.catch((e) => {
				goto('/user/login');
				console.log('Error:' + e);
			})
			.finally(() => {
				loadSession = true;
			});
	});
</script>

{#if loadSession}
	<div class="relative h-screen overflow-y-auto md:flex" data-dev-hint="container">
		{#if $loggedIn}
			<Sidebar />
		{/if}

		<main class="flex-1" class:md:ml-64={$loggedIn}>
			<slot />
		</main>
	</div>
{/if}

<Modal />
<Notification />
