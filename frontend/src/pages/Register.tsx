import { Component, createSignal } from "solid-js";
import "../styles/styles.scss"

export const Register: Component = () => {
    const [getUname, setUname] = createSignal("");
    const [getPswd, setPswd] = createSignal("");
    return (
        <div class="reg">
            <div class="reg-main">
                <div class="reg-title">
                    <div class="reg-title-name">Naarad</div>
                    <div class="reg-title-desc">Register with a unique username and strong password</div>
                </div>
                <div class="reg-input">
                    <div class="reg-uname">
                        <span>Username: </span>
                        <input type="text" onChange={(e) => setUname(e.target.value)} placeholder="Enter a unique username"/>
                    </div>
                    <div class="reg-pswd">
                        <span>Password: </span>
                        <input type="password" onChange={(e) => setPswd(e.target.value)} placeholder="Enter a secure password"/>
                    </div>
                </div>
                <div class="reg-submit">
                    <button class="btn-reg">Register</button>
                </div>
            </div>
        </div>
    )
}