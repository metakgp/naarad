import { Component, createSignal } from "solid-js";
import "../styles/Register.scss"

export const Register: Component = () => {
    const [getUname, setUname] = createSignal("");
    const [getPswd, setPswd] = createSignal("");
    const [getMsg, setMsg] = createSignal("");
    const handleRegister = () => {
        fetch(import.meta.env.VITE_BACKEND_URL+'/register', {
            method: "POST",
            body: JSON.stringify({"uname":getUname(), "pswd": getPswd()}),
        }).then((data) => {
            if(!data.ok){
                data.text().then((res) => {setMsg(res)})
            }else{
                setMsg("Successfully Created User")
            }
        })
    }
    return (
        <div class="reg">
            <div class="reg-main">
                <div class="reg-title">
                    <div class="reg-title-name">Naarad</div>
                    <div class="reg-title-desc">Register with a unique username and strong password.</div>
                </div>
                <div class="reg-input">
                    <div class="reg-uname">
                        <input type="text" onChange={(e) => setUname(e.target.value)} placeholder="Enter a unique username"/>
                    </div>
                    <div class="reg-pswd">
                        <input type="password" onChange={(e) => setPswd(e.target.value)} placeholder="Enter a secure password"/>
                    </div>
                </div>
                <div class="reg-msg">
                    {getMsg()}
                </div>
                <div class="reg-submit">
                    <button class="btn-reg" onClick={() => handleRegister()}>Register</button>
                </div>
                <div class="reg-help">
                    For instructions checkout <a href="./help">Instructions Page</a>.
                </div>
            </div>
        </div>
    )
}