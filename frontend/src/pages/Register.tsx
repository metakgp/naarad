import { Component, createSignal, onMount } from "solid-js";
import "../styles/Register.scss"

export const Register: Component = () => {
    const [getUname, setUname] = createSignal("");
    const [getPswd, setPswd] = createSignal("");
    const [getMsg, setMsg] = createSignal("");

    const [getLenChk, setLenChk] = createSignal(false)
    const [getLwrChk, setLwrChk] = createSignal(false)
    const [getUprChk, setUprChk] = createSignal(false)
    const [getNumChk, setNumChk] = createSignal(false)

    const [pswdChk, setPswdChk] = createSignal(true)
    
    onMount(async () => {
        fetch(import.meta.env.VITE_BACKEND_URL+'/uname', {
            method:"GET",
            credentials: 'include'
        }).then((data) => {
            if(data.ok){
                data.json().then((dataJson) => {
                    var uname = dataJson.email.replace("@kgpian.iitkgp.ac.in", "")
                    setUname(uname)
                })
            }else{
                document.location = "https://heimdall.metakgp.org/"
            }
        })
    })

    const pswdChange = (pswd: string) => {
        
        setLenChk(pswd.length >= 10)
        setLwrChk(/[a-z]/.test(pswd))
        setUprChk(/[A-Z]/.test(pswd))
        setNumChk(/[0-9]/.test(pswd))
        setPswd(pswd);
        if (getLenChk() && getLwrChk() && getUprChk() && getNumChk()){
            setPswdChk(false)
        }else{
            setPswdChk(true)
        }
    }

    const handleRegister = () => {
        fetch(import.meta.env.VITE_BACKEND_URL+'/register', {
            method: "POST",
            body: JSON.stringify({"uname":getUname(), "pswd": getPswd()}),
            credentials: 'include'
        }).then((data) => {
            if(!data.ok){
                data.text().then((res) => {setMsg(res)})
            }else{
                setMsg("Successfully Created User")
                setTimeout(() => {
                    document.location='./help'
                }, 1000)
            }
        })
    }
    return (
        <div class="reg">
            <div class="reg-main">
                <div class="reg-title">
                    <div class="reg-title-name">Naarad</div>
                    <div class="reg-title-desc">Register with username and strong password.</div>
                </div>
                <div class="reg-input">
                    <div class="reg-uname">
                        <input type="text" placeholder="Enter a unique username" value={getUname()} disabled/>
                    </div>
                    <div class="reg-pswd">
                        <input type="password" onInput={(e) => pswdChange(e.target.value)} placeholder="Enter a secure password"/>
                    </div>
                </div>
                <div class="pswd-chk">
                    <div class="pswd-chk-title" hidden={!pswdChk()}>Password Must Contain: </div>
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
                    <button class="btn-reg" onClick={() => handleRegister()} disabled={(getUname() == "") || pswdChk() }>Register</button>
                </div>
                <div class="reg-help">
                    For instructions checkout <a href="./help">Instructions Page</a>.
                </div>
            </div>
        </div>
    )
}