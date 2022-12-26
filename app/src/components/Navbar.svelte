<script>
  import { isLogin } from "../stores/auth";
  import NavLink from "./NavLink.svelte";
</script>

{#if $isLogin}
<nav
  class="fixed top-0 z-20 w-full border-gray-200 bg-gray-900 px-2 py-2.5 sm:px-4"
>
  <div class="container mx-auto flex flex-wrap items-center justify-between">
    <a href="#" class="flex items-center">
      <span
        class="self-center whitespace-nowrap text-xl font-semibold text-white"
      >
        3X2 Banking
      </span>
    </a>
    <div class="hidden w-full md:block md:w-auto" id="navbar-default">
      <ul
        class="mt-4 flex flex-col rounded-lg border p-4 md:mt-0 md:flex-row md:space-x-8 md:border-0 md:text-sm md:font-medium"
      >
        <NavLink pagelink="/profile" pagename="Profile" home="{false}" />
        <NavLink pagelink="/login" pagename="Logout" home="{false}" />
      </ul>
    </div>
  </div>
</nav>
{/if}
