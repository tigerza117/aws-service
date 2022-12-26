<!-- Login Page -->
<script>
  import axios from "axios";
  import { navigate, link } from "svelte-navigator";
  import logo from "../assets/logo.png";

  let email = "";
  let password = "";

  function handleClick() {
    const data = { email: email, password: password };
    axios
      .post("/login", data, { withCredentials: true })
      .then((result) => {
        navigate("/profile");
        alert("Login Complete");
      })
      .catch((err) => {
        console.log(err);
        alert("Invalid Username or Password");
      });
  }
</script>

<section class="h-full">
  <div class="container m-auto h-full py-12 px-6">
    <div
      class="g-6 flex h-full flex-wrap items-center justify-center text-gray-800"
    >
      <div class="xl:w-10/12">
        <div class="block rounded-lg bg-white shadow-lg">
          <div class="g-0 lg:flex lg:flex-wrap">
            <div class="px-4 md:px-0 lg:w-6/12">
              <div class="py-6 md:mx-6 md:p-12">
                <div class="text-center">
                  <img class="mx-auto w-48" src="{logo}" alt="logo" />
                </div>
                <div class="mt-4 space-y-4">
                  <h1
                    class="text-xl font-bold leading-tight tracking-tight text-gray-900 dark:text-white md:text-2xl"
                  >
                    Sign in to your account
                  </h1>
                  <form class="space-y-4 md:space-y-6" action="#">
                    <div>
                      <label
                        for="email"
                        class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
                        >Your email</label
                      >
                      <input
                        type="email"
                        name="email"
                        id="email"
                        class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500 sm:text-sm"
                        placeholder="name@company.com"
                        required=""
                      />
                    </div>
                    <div>
                      <label
                        for="password"
                        class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
                        >Password</label
                      >
                      <input
                        type="password"
                        name="password"
                        id="password"
                        placeholder="••••••••"
                        class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-gray-900 focus:border-primary-600 focus:ring-primary-600 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500 sm:text-sm"
                        required=""
                      />
                    </div>
                    <button
                      type="submit"
                      class="w-full rounded-lg bg-primary-600 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-primary-700 focus:outline-none focus:ring-4 focus:ring-primary-300 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
                    >
                      Sign in
                    </button>
                    <p
                      class="text-sm font-light text-gray-500 dark:text-gray-400"
                    >
                      Don’t have an account yet?
                      <a
                        use:link
                        href="/auth/register"
                        class="font-medium text-primary-600 hover:underline dark:text-primary-500"
                      >
                        Sign up
                      </a>
                    </p>
                  </form>
                </div>
              </div>
            </div>
            <div
              class="flex items-center rounded-b-lg lg:w-6/12 lg:rounded-r-lg lg:rounded-bl-none"
              style="background: linear-gradient(to top left, #13003b, #6d2ab6)"
            >
              <div class="px-4 py-6 text-white md:mx-6 md:p-12">
                <h4 class="mb-6 text-xl font-semibold">
                  We are more than just a company
                </h4>
                <p class="text-sm">
                  We are a financial company that values more than just profits.
                  We prioritize ethics, innovation, and community impact. We
                  strive to be a trusted partner to our clients and a supportive
                  employer to our team. Our focus on sustainability and
                  transparency sets us apart in the industry. We are committed
                  to making a meaningful difference beyond just being a
                  business.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>
