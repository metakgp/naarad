import { Component, createSignal, onMount } from "solid-js";
import "../styles/Register.scss"
import { Spinner } from "../components/Spinner";
import check from "../assets/check.png"
import cross from "../assets/cross.png"

export const Register: Component = () => {
    const [getUname, setUname] = createSignal("");
    const [getMsg, setMsg] = createSignal("");
    const [getIsLoad, setIsLoad] = createSignal(true)
    const [getIsErr, setIsErr] = createSignal(false)
    const [getIsDup, setIsDup] = createSignal(false)
    
    
    onMount(() => {
        fetch(import.meta.env.VITE_BACKEND_URL+'/register', {
            method:"GET",
            credentials: 'include'
        }).then((data) => {
            setIsLoad(false);
            if(data.ok) setMsg("Successfully Created User")
            else if(data.status === 409){
                setIsDup(true);
                setMsg("Username and password is present in IITKGP Email")
            }
            else {
                setIsErr(true)
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
                    <div class="reg-status-uname">Registering user: {getUname()}</div>
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