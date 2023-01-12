<script lang="ts">
	import session from '$lib/session';
	import { onMount } from 'svelte';

	onMount(() => {
		session.logout();
	});
</script>

Logout...
