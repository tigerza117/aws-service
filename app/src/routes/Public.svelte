<script>
  import { navigate } from "svelte-navigator";
  import { isLogin } from "../stores/auth";

  $: if ($isLogin) {
    navigate("/");
  }
</script>

{#if !$isLogin}
<slot />
{/if}
