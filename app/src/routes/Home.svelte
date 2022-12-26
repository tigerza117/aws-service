<script>
  import { link } from "svelte-navigator";
</script>

<div
  class="max flex h-screen flex-col items-center justify-center bg-gray-800 pb-40"
>
  <div
    class="flex h-10 w-10 items-center justify-center rounded-full border-2 bg-gray-900 p-60"
  >
    <div class="absolute h-fit w-full">
      <h1 class="truncate text-center text-5xl font-medium italic text-gray-400"
        >Welcome To</h1
      >
      <h1 class="truncate text-center text-9xl font-bold text-gray-400">3X2</h1>
      <h1
        class="h-20 truncate text-center text-6xl font-medium italic text-gray-400"
        >Banking</h1
      >
    </div>
  </div>
  <div class="flex w-1/4 space-x-5 pt-5">
    <div class="w-1/2 text-gray-100">
      <h1 class="text-center text-xl">Have An Account? </h1>
      <h1 class="text-center text-6xl">
        <a
          use:link
          href="login"
          class="font-medium text-blue-600 hover:underline dark:text-blue-500"
        >
          Login
        </a>
      </h1>
    </div>
    <div class="w-1/2 text-gray-100">
      <h1 class="text-center text-xl">Become Our Member! </h1>
      <h1 class="text-center text-6xl">
        <a
          use:link
          href="register"
          class="font-medium text-blue-600 hover:underline dark:text-blue-500"
          >Register
        </a>
      </h1>
    </div>
  </div>
</div>
