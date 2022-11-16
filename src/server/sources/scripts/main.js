function main(){
    let i
    let HttpReq = new XMLHttpRequest()
    HttpReq.open('GET','/getDates',true)
    HttpReq.send()
    HttpReq.onreadystatechange = ()=> {
        if (HttpReq.readyState === 4 && HttpReq.status=== 200){
            i = JSON.parse(HttpReq.responseText).roomCnt
            console.log(i)
        }
    }
    for (let j = 0; j < i; j++) {
        document.getElementById('box').innerHTML += roomList_meta
    }
}