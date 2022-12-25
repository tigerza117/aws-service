<!-- Register Page -->
<script>
    import axios from "axios";
    import {link, push} from 'svelte-spa-router';
	import NavBarComp from "../components/NavBarComp.svelte";
    import Footer from "../components/FooterComp.svelte";
    let name = "";
    let email = "";
    let password = "";
    
    function handleClick() {
    const data = { name: name, email: email, password: password };
      axios.post("https://i-here-ji.tigerza117.xyz/register", data, {withCredentials: true}).then(result => {
            alert("Register Complete");
            name = '';
			email = '';
			password = '';
            push('/login');
            
		}).catch(err => {
        console.log(err);
      });
    }
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

    <div class="flex flex-col max h-screen bg-gray-800 justify-center items-center pb-40">
		<div class="w-full max-w-sm p-4 rounded-lg shadow-md sm:p-6 md:p-8 bg-gray-700 border-gray-700">
			<form class="space-y-8 w-full">
				<h5 class="text-3xl italic text-center font-medium text-white">Sign In To Our Platform</h5>
                <div>
					<label for="name" class="block mb-2 text-xl font-medium text-white">Enter Your Full Name:</label>
					<input type="name" bind:value={name} name="name" id="name" class="border text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 bg-gray-600 border-gray-500 placeholder-gray-400 text-white" placeholder="John Doe" required>
				</div>
				<div>
					<label for="email" class="block mb-2 text-xl font-medium text-white">Enter Your E-mail:</label>
					<input type="email" bind:value={email} name="email" id="email" class="border text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 bg-gray-600 border-gray-500 placeholder-gray-400 text-white" placeholder="Name@email.com" required>
				</div>
				<div>
					<label for="password"  class="block mb-2 text-xl font-medium text-gray-900 dark:text-white">Enter Your Password:</label>
					<input type="password" bind:value={password} name="password" id="password" placeholder="••••••••••••••••" class=" text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 bg-gray-600 border-gray-500 placeholder-gray-400 text-white" required>
				</div>
				<button on:click={handleClick} class="w-full focus:ring-4 focus:outline-none rounded-lg text-xl font-bold px-5 py-2.5 text-center bg-blue-600 hover:bg-blue-700 focus:ring-blue-800">Register Your Account</button>
				<div class="text-center text-lg font-medium text-gray-300">
					Have An Account? <a href="/login" use:link class="hover:underline text-blue-500">Login Your Account</a>
				</div>
			</form>
		</div>
    </div>

    <Footer/>

</div>