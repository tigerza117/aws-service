<script>
    export let sidename1;
    export let sidename2;
    // import {link} from 'svelte-spa-router';
</script>

<li>
    <div class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white bg-gray-100 dark:bg-gray-700 ">
       <span class="px-2 py-2 text-gray-500 font-semibold text-2xl">{sidename1}: </span>
       <span class="px-2 py-2 text-2xl">{sidename2}</span>

    </div>
 </li>