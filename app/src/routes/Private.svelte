<script>
  import { navigate } from "svelte-navigator";
  import { isLogin } from "../stores/auth";

  $: if (!$isLogin) {
    navigate("/auth/login");
  }
</script>

{#if $isLogin}
<slot />
{/if}
