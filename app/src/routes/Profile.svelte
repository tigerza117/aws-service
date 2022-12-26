<!-- Profile & Account Page -->
<script>
  import SideBarComp from "../components/SideBarComp.svelte";
  import axios from "axios";
  import { onMount } from "svelte";
  import Deposit from "../components/Deposit.svelte";
  import Transfer from "../components/Transfer.svelte";

  let name = "";
  let new_account = "";
  let email = "";
  let userid = "";
  let idcreated_at = "";
  let idupdated_at = "";
  let bank_accounts = [
    {
      name: "Personal",
      no: "1212312312312",
      balance: 10000
    },
    {
      name: "Saving",
      no: "3429873423423",
      balance: 10000
    }
  ];

  axios
    .get("/profile")
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
      .get("/accounts")
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
      .put("/account", data)
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

  function createHex() {
    var hexCode1 = "";
    var hexValues1 = "0123456789abcdef";

    for (var i = 0; i < 6; i++) {
      hexCode1 += hexValues1.charAt(
        Math.floor(Math.random() * hexValues1.length)
      );
    }
    return hexCode1;
  }

  function gen() {
    const deg = Math.floor(Math.random() * 360);
    return "linear-gradient(" +
      deg +
      "deg, " +
      "#" +
      createHex() +
      ", " +
      "#" +
      createHex() +
      ")"
  }

  let  avatarUrl = ""

  onMount(() => {
   fetch("https://randomuser.me/api/")
      .then((results) => {
        return results.json();
      }).then(u => avatarUrl = u.results[0].picture.large);


  })
</script>
<section class="h-full pt-32 pb-24 px-4 md:px-0">
  <div class="max-w-2xl mx-auto h-full">
    <div class="">
      <section class="rounded bg-white p-4 dark:bg-gray-900">
        <div>
          <div
            class="relative mb-16 h-48 w-full rounded"
            style="background: {gen()}"
          >
            <div class="absolute top-[50%] md:left-10">
              <img
                src="{avatarUrl}"
                class="w-48 rounded-full border-8 border-white"
              />
            </div>
          </div>
        </div>
        <hr class="mt-28 my-6"/>
        <div class="px-4">
          <h2 class="mb-4 text-xl font-bold text-gray-900 dark:text-white">
            Profile
          </h2>

          <div action="#">
            <div class="mb-4 grid gap-4 sm:mb-5 sm:grid-cols-2 sm:gap-6">
              <div class="w-full">
                <label
                  for="firstName"
                  class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
                >
                  First name
                </label>
                <input
                  type="text"
                  name="name"
                  id="firstName"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-primary-600 focus:ring-primary-600 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-primary-500 dark:focus:ring-primary-500"
                  value="Apple iMac 27&ldquo;"
                  placeholder="John"
                  required=""
                  disabled
                />
              </div>
              <div class="w-full">
                <label
                  for="lastName"
                  class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
                >
                  First name
                </label>
                <input
                  type="text"
                  name="name"
                  id="lastName"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-primary-600 focus:ring-primary-600 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-primary-500 dark:focus:ring-primary-500"
                  value="Apple iMac 27&ldquo;"
                  placeholder="John"
                  required=""
                  disabled
                />
              </div>
              <div class="w-full">
                <label
                  for="email"
                  class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
                >
                  Email
                </label>
                <input
                  type="email"
                  name="email"
                  id="email"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-primary-600 focus:ring-primary-600 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-primary-500 dark:focus:ring-primary-500"
                  value="Apple"
                  placeholder="Doe"
                  required=""
                  disabled
                />
              </div>
              <div class="w-full">
                <label
                  for="email"
                  class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
                >
                  Email
                </label>
                <input
                  type="email"
                  name="email"
                  id="email"
                  class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-primary-600 focus:ring-primary-600 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-primary-500 dark:focus:ring-primary-500"
                  value="Apple"
                  placeholder="Doe"
                  required=""
                  disabled
                />
              </div>
            </div>
          </div>
        </div>
        <hr class="my-6"/>
        <div class="px-4">
          <div class="justify-between flex">
            <h2 class="mb-4 text-xl font-bold text-gray-900 dark:text-white">
              Accounts
            </h2>
            <div>
              <button class="p-1 px-4 rounded bg-primary-600 text-white">
                <span class="material-icons align-middle">add</span> Create Account
              </button>
            </div>
          </div>
          <div class="mb-4 space-y-4 overflow-x-scroll overflow-auto">
            {#each bank_accounts as bank_account}
              <div class="min-w-[24em] w-full p-4 rounded text-white" style="background: {gen()}">
                <div class="flex justify-between">
                  <div class="space-y-2">
                    <div>
                      {bank_account.name}
                    </div>
                   <div class="flex items-center ">
                     <div>
                       {bank_account.no}
                     </div>
                     <button class="px-2 ">
                       <span class="material-icons align-middle" style="font-size: 1em">content_copy</span>
                     </button>
                   </div>
                  </div>
                  <div >
                   <div class="text-3xl font-bold">
                     {bank_account.balance.toLocaleString('th-TH', {
                       style: 'currency',
                       currency: 'THB',
                     })}
                   </div>
                    <div class="flex justify-end">
                      <Transfer />
                      <Deposit />
                    </div>
                  </div>
                </div>
              </div>
            {/each}
          </div>

        </div>

      </section>
    </div>
  </div>
</section>
