<script setup>
	import InputText from "primevue/inputtext";
	import Password from "primevue/password";
	import { ref } from "vue";
	import { useRouter } from "vue-router";
	import { jwtDecode } from "jwt-decode";
	import axios from "axios";

	const router = useRouter();

	const Email = ref("");
	const Passphrase = ref("");
	const Username = ref("");
	const NoTelp = ref("");

	const signUp = async () => {
		try {
			const res = await axios.post("http://localhost:8080/onize/signup", {
				name: Username.value,
				passphrase: Passphrase.value,
				email: Email.value,
				no_telp: NoTelp.value,
			});

			const token = res.data.token;
			const decoded = jwtDecode(token);
			const UserID = decoded.id;

			localStorage.setItem("token", token);
			localStorage.setItem("userID", UserID);

			router.push("/");
		} catch (err) {
			console.log(err);
		}
	};
</script>
<template>
	<div class="container-center">
		<div class="container" id="container">
			<!-- <div class="form-container log-in-container"> -->
			<form @submit.prevent="signUp">
				<h1 class="login-header">Signup</h1>
				<!-- <input type="email" placeholder="Email" /> -->
				<!-- <input type="password" placeholder="Password" /> -->
				<div class="form-input-container">
					<div class="left-side">
						<div class="form-input">
							<InputText
								id="username"
								v-model="Username"
								aria-describedby="username-help"
								placeholder="Username"
							/>
						</div>
						<div class="form-input">
							<InputText
								id="email"
								v-model="Email"
								aria-describedby="email-help"
								placeholder="Email"
							/>
						</div>
					</div>

					<div class="left-side">
						<div class="form-input">
							<InputText
								id="no-telp"
								v-model="NoTelp"
								aria-describedby="username-help"
								placeholder="No Tlpn"
							/>
						</div>
						<div class="form-input">
							<Password
								toggleMask
								v-model="Passphrase"
								aria-describedby="username-help"
								placeholder="Passphrase"
							/>
						</div>
					</div>
				</div>

				<a href="/login">have an account? <span>Login</span></a>
				<button input="submit">Signup</button>
			</form>
			<!-- </div> -->
			<!-- <div class="overlay-container"> -->
			<!--   <div class="overlay"> -->
			<!--     <div class="overlay-panel overlay-right"> -->
			<!--       <h1>Onize</h1> -->
			<!--       <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Vitae, obcaecati nemo quidem libero impedit dicta -->
			<!--         sapiente aspernatur.</p> -->
			<!--     </div> -->
			<!--   </div> -->
			<!-- </div> -->
		</div>
	</div>
</template>

<style scoped>
	.form-input-container {
		display: flex;
	}
	.form-input {
		margin: 10px;
	}

	.login-header h1 {
		margin-bottom: 100px;
	}

	.container-center {
		width: 100vw;
		height: 100vh;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	h1 {
		font-weight: 450;
	}

	p {
		font-size: 14px;
		font-weight: 100;
		line-height: 20px;
		letter-spacing: 0.5px;
		margin: 20px 0 30px;
	}

	span {
		font-size: 12px;
		color: #6427aa;
	}

	a {
		color: #333;
		font-size: 14px;
		text-decoration: none;
		margin: 15px 0;
	}

	button {
		border-radius: 7px;
		border: 1px solid #6427aa;
		background-color: #6427aa;
		color: #ffffff;
		font-size: 12px;
		font-weight: 400;
		padding: 12px 45px;
		letter-spacing: 1px;
		text-transform: uppercase;
		transition: transform 80ms ease-in;
		width: 100%;
	}

	form {
		background-color: #0b0a0b;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		padding: 0 50px;
		height: 100%;
		text-align: center;
	}

	/* input { */
	/*   background-color: #eee; */
	/*   border: none; */
	/*   padding: 12px 15px; */
	/*   margin: 8px 0; */
	/*   width: 100%; */
	/* } */

	.container {
		/* background-color: #fff; */
		border-radius: 10px;
		box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
		position: relative;
		overflow: hidden;
		width: 768px;
		max-width: 100%;
		min-height: 480px;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.form-container {
		position: absolute;
		top: 0;
		height: 100%;
	}

	.log-in-container {
		left: 0;
		width: 50%;
		z-index: 2;
	}

	.overlay-container {
		position: absolute;
		top: 0;
		left: 50%;
		width: 50%;
		height: 100%;
	}

	.overlay {
		background: #000000;
		background: -webkit-linear-gradient(to right, #6427aa, #340f5e);
		background: linear-gradient(to right, #6427aa, #340f5e);
		background-repeat: no-repeat;
		background-size: cover;
		background-position: 0 0;
		color: #ffffff;
		position: relative;
		left: -100%;
		height: 100%;
		width: 200%;
	}

	.overlay-panel {
		position: absolute;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		padding: 0 40px;
		text-align: center;
		top: 0;
		height: 100%;
		width: 50%;
	}

	.overlay-right {
		right: 0;
	}
</style>
