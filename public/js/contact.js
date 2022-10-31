const button = document.querySelector('.submit button');

button.addEventListener('click', function() {
    const nama = document.querySelector('.nama');
    const email = document.querySelector('.email');
    const phone = document.querySelector('.phone');
    const subject = document.querySelector('.subject');
    const message = document.querySelector('.message');
    const emailReceiver = "faequl96@gmail.com";

    
    function ObjData() {
        this.nama = nama.value;
        this.email = email.value;
        this.phone = phone.value;
        this.subject = subject.value;
        this.message = message.value;
        this.emailReceiver = emailReceiver;
    }
    const dataForm = new ObjData();

    if(dataForm.nama != "" && dataForm.email != "" && dataForm.phone != "" && dataForm.subject != "" && dataForm.message != "") {
        
        console.log(dataForm);

        let a = document.createElement('a');
        a.href = `https://mail.google.com/mail/?view=cm&fs=1&to=${dataForm.emailReceiver}&su=${dataForm.subject}&body=Halo, nama saya ${dataForm.nama}, ${dataForm.message}, silahkan hubungi saya ke nomor ${dataForm.phone} atau email ke ${dataForm.email}`;
        a.target = '_blank'
        a.click()

        nama.value = "";
        email.value = "";
        phone.value = "";
        subject.value = "";
        message.value = "";

        button.style.backgroundColor = '#e4e4e4';
        button.style.color = '#b2abab';
        button.style.cursor = 'default';

    } else {
        alert('Semua field wajib diisi');
    }

})