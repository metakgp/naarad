import { Component, createSignal, onMount } from "solid-js";
import { BACKEND_URL } from "../constants";
import { Spinner } from "../components/Spinner";
import check from "../assets/check.png"
import cross from "../assets/cross.png"

export const Register: Component = () => {
    const [getStatus, setStatus] = createSignal("Registering user...");
    const [getMsg, setMsg] = createSignal("");
    const [getIsLoad, setIsLoad] = createSignal(true)
    const [getIsErr, setIsErr] = createSignal(false)
    const [getIsDup, setIsDup] = createSignal(false)
    
    
    onMount(() => {
        fetch(BACKEND_URL+'/register', {
            method:"GET",
            credentials: 'include'
        }).then((data) => {
            setIsLoad(false);
            if(data.ok){
                setStatus("User Registration Successful!")
                setMsg("Check your institute email for login credentials")

                setTimeout(() => {
                    document.location = "https://naarad.metakgp.org/login"
                }, 3000);
            }
            else if(data.status === 409){
                setIsDup(true);
                setStatus("User Already Registered!")
                setMsg("Check your institute mail for login credentials")

                setTimeout(() => {
                    document.location = "https://naarad.metakgp.org/login"
                }, 3000);
            }
            else if(data.status == 401){
                document.location = "https://heimdall.metakgp.org?redirect_url=https://naarad-signup.metakgp.org"
            }
            else {
                setIsErr(true)
                setStatus("User Registration Error!")
                data.text().then((bodyData) => {
                    setMsg(bodyData)
                })
            }
        })
    })

    return (
        <div class="reg">
            <div class="reg-main">
                <div class="reg-title">
                    <div class="reg-title-name">
                        MetaKGP Naarad
                    </div>
                    <div class="reg-title-desc">
                        Naarad Registration for accessing notifications
                    </div>
                </div>
                <div class="reg-status">
                    <div class="reg-status-title">{getStatus()}</div>
                    <div class="reg-status-svg">
                        {getIsLoad() == true ? <Spinner /> : (getIsDup() == true ? <img src={cross} /> : (getIsErr() == true ? <img src={cross}/> : <img src={check} />))}
                    </div>
                    <div class="reg-status-text">{getMsg()}</div>
                </div>
                <div class="reg-footer">
                    <h3 class="reg-footer">Made with ❤️ and {"</>"} by <a href="https://github.com/metakgp/naarad" target="_blank">MetaKGP</a></h3>
                </div>
            </div>
        </div>
    )
}