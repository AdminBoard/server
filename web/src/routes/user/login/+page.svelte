<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import notification from '$lib/components/notification/notification';
	import Textfield from '$lib/components/textfield.svelte';
	import session from '$lib/session';
	import { get } from 'svelte/store';

	var username = '',
		password = '';
	var disabled = false;

	function login() {
		disabled = true;
		session
			.login(username, password)
			.then(() => {
				const redirect = get(page).url.searchParams.get('redirect');
				if (redirect == null || redirect.trim() == '')
					goto('/dashboard');
				else goto(redirect);
			})
			.catch((e) => notification.show('Login Error', e))
			.finally(() => (disabled = false));
	}
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-100">
	<div class="px-8 py-6 mt-4 text-left bg-white shadow-lg">
		<div class="flex justify-center">
			<img src="/logo.jpg" alt="logo" />
		</div>
		<form action="">
			<div class="mt-4 space-y-1">
				<Textfield label="Username" bind:value={username} {disabled} />
				<Textfield
					label="Password"
					placeholder=""
					type="password"
					bind:value={password}
					{disabled}
				/>
				<div>
					<button class="mt-6 float-right" {disabled} on:click={login}
						>Login</button
					>
				</div>
			</div>
		</form>
	</div>
</div>
