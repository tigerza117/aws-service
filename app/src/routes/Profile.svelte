<!-- Profile & Account Page -->
<script>
    import NavBarComp from "../components/NavBarComp.svelte";
    import SideBarComp from "../components/SideBarComp.svelte";
    import Footer from "../components/FooterComp.svelte";
    import axios from "axios";
    import { link } from "svelte-spa-router";
    let name = "";
    let new_account = "";
    let email = "";
    let userid = "";
    let idcreated_at = "";
    let idupdated_at = "";
    let bank_accounts = [];

    axios
        .get("https://i-here-ji.tigerza117.xyz/profile")
        .then(function (response) {
            if (response.data) {
                console.log(JSON.stringify(response.data));
                console.log(response.data.name);
                name = response.data.name;
                email = response.data.email;
                userid = response.data.id;
                idcreated_at = response.data.created_at;
                idupdated_at = response.data.updated_at;
                bank_accounts = response.data.account;
            } else {
                console.log("No data received from the server");
            }
        })
        .catch(function (error) {
            console.log(error);
        });

    function reload() {
        axios
            .get("https://i-here-ji.tigerza117.xyz/accounts")
            .then(function (response) {
                if (response.data) {
                    console.log(JSON.stringify(response.data));
                    console.log(response.data.name);
                    bank_accounts = response.data;
                } else {
                    console.log("No data received from the server");
                }
            })
            .catch(function (error) {
                console.log(error);
            });
    }
    reload();

    // /*Create Bank Account*/
    function handleClick() {
        const data = { name: new_account };
        axios
            .put("https://i-here-ji.tigerza117.xyz/account", data)
            .then((result) => {
                new_account = "";
                reload();
            })
            .catch((err) => {
                console.log(err);
                alert("Invalid Name, please input Another Unique Account Name");
                new_account = "";
            });
    }
</script>


<div style="overflow: hidden; background-color: black; max-width:100%; width:100%; max-height:100vh;">
        
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
                    pagelink="/profile" 
                    pagename="Profile" 
                    home={true} 
                />
                    <NavBarComp
                        pagelink="/transfer"
                        pagename="Transfer"
                        home={false}
                    />
                    <NavBarComp
                        pagelink="/deposit"
                        pagename="Deposit"
                        home={false}
                    />
                    <NavBarComp
                        pagelink="/login"
                        pagename="Logout"
                        home={false}
                    />
                </ul>
            </div>
        </div>
    </nav>

    <div class="flex max h-auto bg-gray-500">
        <div class="w-1/4 h-1/4">
            <div class="overflow-y-auto py-4 px-3 bg-gray-800 h-screen">
                <div style="justify-content: center; align-item:center; display:flex; padding: 3%; ">
                    <img
                        style="width: 200px; height: 200px; border-radius:100%; background-color:azure;"
                        src="https://static.thenounproject.com/png/4738596-200.png"
                        alt="profile-icon"
                    />
                </div>
                <ul class="space-y-2">
                    <SideBarComp sidename1="Name" sidename2={name} />
                    <SideBarComp sidename1="Email" sidename2={email} />
                    <SideBarComp sidename1="User-ID" sidename2={userid} /><br>
                    <h4 class="text-center text-zinc-100">ID Created: {idcreated_at}</h4>
                    <h4 class="text-center text-zinc-100">ID Updated: {idupdated_at}</h4>
                </ul>
            </div>

        </div>
        <div class="w-3/4 bg-gray-800">
            <div class="">

                    <h1 class="mb-2 text-center text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Create New Bank Account</h1>
                    <input
                    class="border text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 bg-gray-600 border-gray-500 text-white"
                        bind:value={new_account}
                        type="new_account"
                        name="new_account"
                        placeholder=" Account Name"
                    />
                    <button class="text-white w-full focus:ring-4 focus:outline-none font-medium rounded-lg text-sm px-5 py-2.5 text-center bg-blue-600 hover:bg-blue-700 focus:ring-blue-800" on:click={handleClick}>Create</button>
            </div>
            <div class="">
                <div class="flow-root overflow-y-auto h-[41.5rem] w-full">
                    {#each bank_accounts as bank_account}
                    <ul class="divide-y divide-gray-700 bg-gray-700" style="padding-left: 1%; border: white; border-style:groove;">
                        <li class="py-3 sm:py-4">
                            <div class="flex items-center space-x-4">
                                <div class="flex-1 min-w-0 w-3/6">
                                    <div class="flex">
                                        <div class="pr-2">
                                            <p class="text-6xl font-medium truncate text-white">
                                            {bank_account.name}
                                            </p>
                                        </div>
                                        <div class="flex justify-center items-end">
                                            <p class="text-2xl truncate text-gray-400">ID: {bank_account.id}</p>
                                        </div>
                                    </div>
                                    <br>
                                    <p class="text-3xl truncate text-gray-400">
                                        No.{bank_account.no} 
                                    </p>
                                    <br><p class="text-sm truncate text-gray-400">
                                        Account Created: {bank_account.created_at}
                                    </p>
                                    <p class="text-sm truncate text-gray-400">
                                        Account Updated: {bank_account.updated_at}
                                    </p>
                                </div>
                                <div class="pr-5 column-flex justify-end text-base font-semibold text-white w-3/6">
                                    <h1 class="text-4xl text-right">Balance: </h1>
                                    <h1 class="text-right text-8xl">${bank_account.balance}</h1>
                                </div>
                            </div>
                        </li>
                    </ul>
                    {/each}
                </div>

            </div>
        </div>
    </div>

    <Footer/>
</div>