import { Component, createSignal } from "solid-js";
import "../styles/Register.scss"

export const Register: Component = () => {
    const [getUname, setUname] = createSignal("");
    const [getPswd, setPswd] = createSignal("");
    const [getMsg, setMsg] = createSignal("");

    const [getLenChk, setLenChk] = createSignal(false)
    const [getLwrChk, setLwrChk] = createSignal(false)
    const [getUprChk, setUprChk] = createSignal(false)
    const [getNumChk, setNumChk] = createSignal(false)

    const [btnDis, setBtnDis] = createSignal(true)
    

    const pswdChange = (pswd: string) => {
        
        setLenChk(pswd.length >= 10)
        setLwrChk(/[a-z]/.test(pswd))
        setUprChk(/[A-Z]/.test(pswd))
        setNumChk(/[0-9]/.test(pswd))
        setPswd(pswd);
        if (getLenChk() && getLwrChk() && getUprChk() && getNumChk()){
            setBtnDis(false)
        }else{
            setBtnDis(true)
        }
    }

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
                        <input type="text" onInput={(e) => setUname(e.target.value)} placeholder="Enter a unique username"/>
                    </div>
                    <div class="reg-pswd">
                        <input type="password" onInput={(e) => pswdChange(e.target.value)} placeholder="Enter a secure password"/>
                    </div>
                </div>
                <div class="pswd-chk">
                    <div class="pswd-chk-title" hidden={!btnDis()}>Password Must Contain: </div>
                    <ul>
                        <li hidden={getLenChk()}>
                            Minimum 10 characters long
                        </li>
                        <li hidden={getLwrChk()}>
                            At least 1 lowercase letter
                        </li>
                        <li hidden={getUprChk()}>
                            At least 1 uppercase letter
                        </li>
                        <li hidden={getNumChk()}>
                            At least 1 number
                        </li>
                    </ul>
                </div>
                <div class="reg-msg">
                    {getMsg()}
                </div>
                <div class="reg-submit">
                    <button class="btn-reg" onClick={() => handleRegister()} disabled={btnDis()}>Register</button>
                </div>
                <div class="reg-help">
                    For instructions checkout <a href="./help">Instructions Page</a>.
                </div>
            </div>
        </div>
    )
}