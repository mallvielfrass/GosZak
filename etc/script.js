function a(){
    labeltext1.innerText=""
    var library = {};
    var myForm = document.forms["forms"]; 
    i=0;
    for (var key in myForm.elements["interest"]) {
        try {
            if  ( myForm.elements["interest"][i].checked==true)  {
            console.log(key, myForm.elements["interest"][i].id, myForm.elements["interest"][i].checked);
            library[myForm.elements["interest"][i].id]= myForm.elements["interest"][i].value;
        }
          } catch (err) {
            console.log("error")// обработка ошибки 
            break;
          } 
      //  console.log(key, myForm.elements["interest"][i].value);
        i=i+1; 
    }
   
    
  if(myForm.elements["sort1"].value!=""){
    library["sortDirection"]= myForm.elements["sort1"].value;
  }
  if(myForm.elements["radiobutton2"].value!=""){
    library["sortBy"]= myForm.elements["radiobutton2"].value;
  }
  if(myForm.elements["selectCity"].value!=""){
    library["is_north_west_district"]= myForm.elements["selectCity"].value;
  }
  if(myForm.elements["date"].value!=""){
    var msUTC = Date.parse(myForm.elements["date"].value);
    var now = new Date(msUTC);
        
    library["publishDateFrom"]= now.getDate()+"."+(now.getMonth()+1)+"."+now.getFullYear();
  }
  if(myForm.elements["date2"].value!=""){
    var msUTC = Date.parse(myForm.elements["date2"].value);
    var now = new Date(msUTC);
        
    library["publishDateTo"]= now.getDate()+"."+(now.getMonth()+1)+"."+now.getFullYear();
  }
  
   //console.log( myForm.elements["sort1"][0].value)
  // console.log(document.forms["forms"].elements["interest"][1].value);
  console.log("key start")
  var labellibrary="";
  for (var key in library){
    if (library.hasOwnProperty(key)) {

        console.log(key + " -> " + library[key]);
        labellibrary=labellibrary+"&"+key + "=" + library[key];
    }
  }
 // labeltext1.innerText=labellibrary;
  send(labellibrary)
  //  labeltext1.innerText="lol";
    //console.log("hi")
}
async function send(xlibrary) {

    let response = await fetch('http://localhost:8080/api',{
        method: "POST",
        body: xlibrary
    });

    document.querySelector('#labeltext1').innerHTML = await response.text();
}