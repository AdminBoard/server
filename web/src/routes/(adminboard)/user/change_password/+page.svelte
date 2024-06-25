<script>
	import notification from '$lib/components/notification/notification';
	import Textfield from '$lib/components/textfield.svelte';
	import session from '$lib/session';
	import { get } from 'svelte/store';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	let disabled = false;
	let password = '',
		confirm = '';

	function click() {
		disabled = true;
		session
			.changePassword(password, confirm)
			.then((ok) => {
				if (ok) {
					const redirect = get(page).url.searchParams.get('redirect');
					if (redirect == null || redirect.trim() == '')
						goto('/dashboard');
					else goto(redirect);
				}
			})
			.catch((e) => notification.show('Error', e))
			.finally(() => (disabled = false));
	}
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-100">
	<div class="px-8 py-6 mt-4 text-left bg-white shadow-lg">
		<div class="flex justify-center">
			<img src="/logo.jpg" alt="logo" />
		</div>
		<div class="mt-4 space-y-1">
			<Textfield
				label="New Password"
				type="password"
				bind:value={password}
				{disabled}
			/>
			<Textfield
				label="Confirm Password"
				type="password"
				bind:value={confirm}
				{disabled}
			/>
			<div>
				<button class="mt-6 float-right" {disabled} on:click={click}
					>Change Password</button
				>
			</div>
		</div>
	</div>
</div>
