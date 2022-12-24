<script>
    import NavBarComp from "../components/NavBarComp.svelte";
    import SideBarComp from "../components/SideBarComp.svelte";
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
                alert("Create Complete");
                new_account = "";
                reload();
            })
            .catch((err) => {
                console.log(err);
                alert(err);
            });
    }
</script>


<div style="overflow: hidden; background-color: black; max-width:100%; width:100%; max-height:100vh;">
        
    <!-- navbar -->
    <nav class="w-full h-[10%] bg-white border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-900">
        <div class="container flex flex-wrap items-center justify-between mx-auto">
            <a href="https://flowbite.com/" class="flex items-center">
                <img
                    src="https://flowbite.com/docs/images/logo.svg"
                    class="h-6 mr-3 sm:h-9"
                    alt="Flowbite Logo"
                />
                <span
                    class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
                    >3X2</span
                >
            </a>
            <div class="hidden w-full md:block md:w-auto" id="navbar-default">
                <ul
                    class="flex flex-col p-4 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700"
                >
                    <NavBarComp 
                        pagelink="/" 
                        pagename="Home" 
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
    <!-- navbar dark:bg-gray-900 -->

    <!-- <a href="/transfer" use:link>Transfer</a> -->
    <!-- <a href="/deposit" use:link>Deposit</a> -->
    <!-- <a href="/login" use:link>Logout</a> -->
    <div class="flex max h-auto dark:bg-gray-500">
        <div class="w-1/4 h-1/4">
            <div class="overflow-y-auto py-4 px-3 bg-gray-50 rounded dark:bg-gray-800 h-screen">
                <!-- <h1 class="text-center font-semibold dark:text-white text-5xl">
                    INFORMATION
                </h1> -->
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
            <!-- <h2 class="text-2xl">Name: {name}</h2>
            <h2>Email: {email}</h2>
            <h2>UserID: {userid}</h2>
            <h4>ID Created: {idcreated_at}</h4>
            <h4>ID Updated: {idupdated_at}</h4> -->
        </div>
        <div class="w-3/4">
            <div class="">
                <section>
                    <h1>Create New Bank Account</h1>
                    <input
                        bind:value={new_account}
                        type="new_account"
                        name="new_account"
                        placeholder=" Account Name"
                    />
                    <button on:click={handleClick}>Create</button>
                </section>
            </div>
            <div class="">
                <div class="flow-root overflow-y-auto h-[42rem] w-full rounded-lg">
                    {#each bank_accounts as bank_account}
                    <ul class="divide-y divide-gray-200 dark:divide-gray-700 bg-gray-700" style="padding: 1%; border: white; border-style:groove;">
                        <li class="py-3 sm:py-4">
                            <div class="flex items-center space-x-4">
                                <div class="flex-1 min-w-0 w-4/6">
                                    <p class="text-4xl font-medium text-gray-900 truncate dark:text-white">
                                        {bank_account.name}
                                    </p>
                                    <br>
                                    <div class="h-full flex">
                                        <div class="w-1/2">
                                            <p class="text-xl text-gray-500 truncate dark:text-gray-400">
                                                No.{bank_account.no} 
                                            </p>
                                        </div>
                                        <div class="w-1/2">
                                            <p class="text-xl text-gray-500 truncate dark:text-gray-400">
                                                ID: {bank_account.id}
                                            </p>
                                        </div>
                                    </div>
                                    <br><p class="text-sm text-gray-500 truncate dark:text-gray-400">
                                        Account Created: {bank_account.created_at}
                                    </p>
                                    <p class="text-sm text-gray-500 truncate dark:text-gray-400">
                                        Account Updated: {bank_account.updated_at}
                                    </p>
                                </div>
                                <div class="column-flex items-center text-base font-semibold text-gray-900 dark:text-white w-2/6">
                                    <h1 class="text-center text-4xl">Balance</h1>
                                    <h1 class="text-center text-8xl">${bank_account.balance}</h1>
                                </div>
                            </div>
                        </li>
                    </ul>
                    {/each}
                </div>
                <!-- <div class="overflow-y-auto h-96 w-full">
                    {#each bank_accounts as bank_account}
                        <div>
                        <div class="ba-card" style="border-style:dotted; width:100%">
                            <!-- <p>Account</p> -->
                            <!-- <h1>{bank_account.name}</h1>
                            <h3>No.{bank_account.no}</h3>
                            <h3>ID: {bank_account.id}</h3>
                            <h2>Balance: {bank_account.balance}</h2>
                            <h4>Account Created: {bank_account.created_at}</h4>
                            <h4>Account Updated: {bank_account.updated_at}</h4>
                        </div>
                        <div class="ba-card" style="border-style:dotted; width:100%">
                            <h2>Balance: {bank_account.balance}</h2>
                        </div>
                        </div>
                    {/each}
                </div> -->
            </div>
        </div>
    </div>

    <!-- footer -->
    <footer
        class="fixed h-[10%] bottom-0 left-0 z-20 p-4 w-full bg-white border-t border-gray-200 shadow md:flex md:items-center md:justify-between md:p-6 dark:bg-gray-800 dark:border-gray-600"
    >
        <span class="text-sm text-gray-500 sm:text-center dark:text-gray-400"
            >© 2022 <a href="https://flowbite.com/" class="hover:underline"
                >Flowbite™</a
            >. All Rights Reserved.
        </span>
        <ul
            class="flex flex-wrap items-center mt-3 text-sm text-gray-500 dark:text-gray-400 sm:mt-0"
        />
    </footer>
    <!-- footer -->
</div>

<style>

    section {
        text-align: center;
        justify-content: center;
        width: 100%;
        display: grid;
        grid-template-columns: 1fr 1fr;
        padding: 10px;
        box-shadow: 2px 2px 4px #dedede;
        border: 1px solid #888;
        margin-bottom: 10px;
        display: flex;
        flex-direction: column;
    }
</style>
