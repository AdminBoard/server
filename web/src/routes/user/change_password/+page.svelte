<script>
	import notification from '$lib/components/notification/notification';
	import Textfield from '$lib/components/textfield.svelte';

	let disabled = false;
	let password1 = '',
		password2 = '';

	function click() {
		if (password1.length < 6) {
			notification.error('Password minimum length is 6 characters');
		} else if (password1 != password2) {
			notification.error('Password not matches');
		} else {
			disabled = true;

			disabled = false;
		}
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
				bind:value={password1}
				{disabled}
			/>
			<Textfield
				label="Repeat Password"
				type="password"
				bind:value={password2}
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
