<script>
    import {link} from 'svelte-spa-router';
    import NavBarComp from "../components/NavBarComp.svelte";
    import Footer from "../components/FooterComp.svelte";
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

    <!-- <div class="flex max h-auto bg-gray-800">
        <div class="rounded-full bg-gray-900 border-2 flex justify-center items-center p-60 w-10 h-10">
            <div class="absolute w-full h-fit">
                <h1 class="text-4xl font-medium text-center truncate text-gray-400">Welcome to</h1>
                <h1 class="text-9xl font-bold text-center truncate text-gray-400">3x2</h1>
                <h1 class="text-5xl font-medium text-center truncate text-gray-400 h-20">Banking</h1>
            </div>
        </div>
        <div>
            <p>Have an account? <a href="/login" use:link>Login</a></p>
            <p>Become our member! <a href="/Register" use:link>Register</a></p>
        </div>
    </div> -->
    <div class="flex flex-col max h-screen bg-gray-800 justify-center items-center pb-40">
        <div class="rounded-full bg-gray-900 border-2 flex justify-center items-center p-60 w-10 h-10">
            <div class="absolute w-full h-fit">
                <h1 class="text-5xl italic font-medium text-center truncate text-gray-400">Welcome To</h1>
                <h1 class="text-9xl font-bold text-center truncate text-gray-400">3X2</h1>
                <h1 class="text-6xl italic font-medium text-center truncate text-gray-400 h-20">Banking</h1>
            </div>
        </div>
        <div class="flex pt-5 w-1/4 space-x-5">
            <div class="w-1/2 text-gray-100">
                <h1 class="text-center text-xl">Have An Account? </h1>
                <h1 class="text-center text-6xl"><a href="/login" use:link class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Login</a></h1>
            </div>
            <div class="w-1/2 text-gray-100">
                <h1 class="text-center text-xl">Become Our Member! </h1>
                <h1 class="text-center text-6xl"><a href="/register" use:link class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Register</a></h1>
            </div>
        </div>
    </div>

    <!-- <footer
        class="fixed h-[5%] bottom-0 left-0 z-20 p-4 w-full shadow md:flex md:items-center md:justify-between md:p-6 bg-gray-800 border-gray-600"
        >
        <span class="text-sm  sm:text-center text-gray-400"
            >© 2022 <a href="https://flowbite.com/" class="hover:underline"
                >3X2™</a
            >. All Rights Reserved.
        </span>
        <ul
            class="flex flex-wrap items-center mt-3 text-sm text-gray-400 sm:mt-0"
        />
    </footer> -->
    <Footer/>
</div>

<style>
    /* body{
        text-align: center;
        justify-content: center;
        height: 100vh;
        width: 100vw;
    } */
</style>