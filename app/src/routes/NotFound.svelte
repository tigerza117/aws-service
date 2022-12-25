<!-- Not-Found Page -->
<script>
    import {link} from 'svelte-spa-router';
    import Footer from "../components/FooterComp.svelte";
    import NavBarComp from '../components/NavBarComp.svelte';
</script>


<div style="overflow: hidden; background-color: black; max-width:100%; width:100%; max-height:100vh;" class="font-medium">

    <!-- navbar -->
    <nav class="w-full h-[10%] border-gray-200 px-2 sm:px-4 py-2.5 rounded bg-gray-900">
        <div class="container flex flex-wrap items-center justify-between mx-auto">
            <a href="https://flowbite.com/" class="flex items-center">
                <img
                    src="https://flowbite.com/docs/images/logo.svg"
                    class="h-6 mr-3 sm:h-9"
                    alt="Flowbite Logo"
                />
                <span
                    class="self-center text-xl font-semibold whitespace-nowrap text-white"
                    >3X2 Banking</span
                >
            </a>
            <div class="hidden w-full md:block md:w-auto" id="navbar-default">
                <ul
                    class="flex flex-col p-4 mt-4 border rounded-lg md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0"
                >
                    <NavBarComp 
                        pagelink="/" 
                        pagename="Home" 
                        home={true} 
                    />
                    <NavBarComp
                    pagelink="/login"
                    pagename="Login"
                    home={false}
                    />
                    <NavBarComp
                        pagelink="/register"
                        pagename="Register"
                        home={false}
                    />
                </ul>
            </div>
        </div>
    </nav>
    <!-- navbar -->

    <div class="flex flex-col max h-screen bg-gray-800 justify-center items-center pb-40">
        <div class="rounded-full bg-gray-900 border-2 flex justify-center items-center p-60 w-10 h-10">
            <div class="absolute w-full h-fit">
                <h1 class="text-5xl italic font-medium text-center truncate text-gray-400">NOT</h1>
                <!-- <h1 class="text-9xl font-bold text-center truncate text-gray-400">3X2</h1> -->
                <h1 class="text-5xl italic font-medium text-center truncate text-gray-400 h-20">FOUND</h1>
            </div>
        </div>
    </div>
    <Footer/>
</div>

