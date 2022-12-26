<script>
  let isOpen = false;
</script>

<button class="p-1 text-xs" on:click={() => isOpen = !isOpen}>Transfer</button>

{#if isOpen}
  <div
    id="defaultModal"
    tabindex="-1"
    aria-hidden="true"
    class="h-modal  fixed top-0 left-0 right-0 z-50 w-full overflow-y-auto overflow-x-hidden p-4 md:inset-0 md:h-full backdrop-blur-sm drop-shadow-2xl"
  >
    <div class="relative h-full w-full max-w-2xl md:h-auto m-auto mt-32 ">
      <!-- Modal content -->
      <div class="relative rounded-lg bg-white shadow dark:bg-gray-700">
        <!-- Modal header -->
        <div
          class="flex items-start justify-between rounded-t border-b p-4 dark:border-gray-600"
        >
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
            Transfer
          </h3>
          <button
            type="button"
            class="ml-auto inline-flex items-center rounded-lg bg-transparent p-1.5 text-sm text-gray-400 hover:bg-gray-200 hover:text-gray-900 dark:hover:bg-gray-600 dark:hover:text-white"
            data-modal-toggle="defaultModal"
            on:click={() => isOpen = false}
          >
            <svg
              aria-hidden="true"
              class="h-5 w-5"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span class="sr-only">Close modal</span>
          </button>
        </div>
        <!-- Modal body -->
        <div class="space-y-6 p-6">
          <div>
            <label for="amount" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Amount</label>
            <input type="number" name="price" id="amount"
                   class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                   placeholder="฿2999" required="">
          </div>
        </div>
        <!-- Modal footer -->
        <div
          class="flex items-center space-x-2 rounded-b border-t border-gray-200 p-6 dark:border-gray-600"
        >
          <button
            data-modal-toggle="defaultModal"
            type="button"
            class="rounded-lg bg-blue-700 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Transfer
          </button>
        </div>
      </div>
    </div>
  </div>

{/if}
