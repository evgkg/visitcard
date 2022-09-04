let canCopy = true;

document.onclick = event => {
    const elem = event.target;
    //.log(elem);
    if(!canCopy){return;}
    if(elem.classList.contains('emailValue')){
        navigator.clipboard.writeText(elem.innerHTML)
            .then(()=>{
                canCopy = false;
                document.querySelector('.emailValue').innerHTML += '<span id="copied">  copied</span>'
                setTimeout(clearText, 1000);
            })
    }
}
function clearText(){
    cp = document.getElementById('copied to clipboard')
    cp.remove();
    canCopy = true;
}
