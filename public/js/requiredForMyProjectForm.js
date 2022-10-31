const projectName = document.querySelector('.projectName');
const startDate = document.querySelector('.startDate');
const endDate = document.querySelector('.endDate');
const description = document.querySelector('.description');
const uploadImg = document.querySelector('.uploadImg');
const renderImg = uploadImg.files;
const create = document.querySelector('.create');


projectName.addEventListener('click', function() {
    
    projectName.classList.add('pin');
    projectName.removeAttribute('placeholder');
    if(startDate.classList.contains('pin') && startDate.value == "") {
        startDate.style.border = '2px solid red';
    }
    if(endDate.classList.contains('pin') && endDate.value == "") {
        endDate.style.border = '2px solid red';
    }
    if(description.classList.contains('pin') && description.value == "") {
        description.style.border = '2px solid red';
    }
    if(uploadImg.classList.contains('pin') && renderImg.length == 0) {
        document.querySelector('.boxField').style.border = '2px solid red';
    }
    projectName.addEventListener('input', function() {
        if(projectName.value == "") {
            projectName.style.border = '2px solid red';
            create.style.backgroundColor = '#e4e4e4';
            create.style.color = '#b2abab';
            create.style.cursor = 'default';
            create.removeAttribute('type');
            create.setAttribute('type', 'button');
        } else{
            projectName.style.border = '2px solid rgba(0,0,0,0)';
            projectName.classList.remove('pin');
            if(projectName.value != "" && startDate.value != "" && endDate.value != "" && description.value != "" && uploadImg.value != "") {
                create.style.backgroundColor = '#000';
                create.style.color = '#eaeaea';
                create.style.cursor = 'pointer';
                create.removeAttribute('type');
                create.setAttribute('type', 'submit');
            }
        }
    })
})


startDate.addEventListener('click', function() {
    
    startDate.classList.add('pin');
    startDate.removeAttribute('placeholder');
    if(projectName.classList.contains('pin') && projectName.value == "") {
        projectName.style.border = '2px solid red';
    }
    if(endDate.classList.contains('pin') && endDate.value == "") {
        endDate.style.border = '2px solid red';
    }
    if(description.classList.contains('pin') && description.value == "") {
        description.style.border = '2px solid red';
    }
    if(uploadImg.classList.contains('pin') && renderImg.length == 0) {
        document.querySelector('.boxField').style.border = '2px solid red';
    }
    startDate.addEventListener('input', function() {
        if(startDate.value == "") {
            startDate.style.border = '2px solid red';
            create.style.backgroundColor = '#e4e4e4';
            create.style.color = '#b2abab';
            create.style.cursor = 'default';
            create.removeAttribute('type');
            create.setAttribute('type', 'button');
        } else{
            startDate.style.border = '2px solid rgba(0,0,0,0)';
            startDate.classList.remove('pin');
            if(projectName.value != "" && startDate.value != "" && endDate.value != "" && description.value != "" && uploadImg.value != "") {
                create.style.backgroundColor = '#000';
                create.style.color = '#eaeaea';
                create.style.cursor = 'pointer';
                create.removeAttribute('type');
                create.setAttribute('type', 'submit');
            }
        }
    })
})


endDate.addEventListener('click', function() {
    
    endDate.classList.add('pin');
    endDate.removeAttribute('placeholder');
    if(startDate.classList.contains('pin') && startDate.value == "") {
        startDate.style.border = '2px solid red';
    }
    if(projectName.classList.contains('pin') && projectName.value == "") {
        projectName.style.border = '2px solid red';
    }
    if(description.classList.contains('pin') && description.value == "") {
        description.style.border = '2px solid red';
    }
    if(uploadImg.classList.contains('pin') && renderImg.length == 0) {
        document.querySelector('.boxField').style.border = '2px solid red';
    }
    endDate.addEventListener('input', function() {
        if(endDate.value == "") {
            endDate.style.border = '2px solid red';
            create.style.backgroundColor = '#e4e4e4';
            create.style.color = '#b2abab';
            create.style.cursor = 'default';
            create.removeAttribute('type');
            create.setAttribute('type', 'button');
        } else{
            endDate.style.border = '2px solid rgba(0,0,0,0)';
            endDate.classList.remove('pin');
            if(projectName.value != "" && startDate.value != "" && endDate.value != "" && description.value != "" && uploadImg.value != "") {
                create.style.backgroundColor = '#000';
                create.style.color = '#eaeaea';
                create.style.cursor = 'pointer';
                create.removeAttribute('type');
                create.setAttribute('type', 'submit');
            }
        }
    })
})


description.addEventListener('click', function() {
    
    description.classList.add('pin');
    description.removeAttribute('placeholder');
    if(startDate.classList.contains('pin') && startDate.value == "") {
        startDate.style.border = '2px solid red';
    }
    if(endDate.classList.contains('pin') && endDate.value == "") {
        endDate.style.border = '2px solid red';
    }
    if(projectName.classList.contains('pin') && projectName.value == "") {
        projectName.style.border = '2px solid red';
    }
    if(uploadImg.classList.contains('pin') && renderImg.length == 0) {
        document.querySelector('.boxField').style.border = '2px solid red';
    }
    description.addEventListener('input', function() {
        if(description.value == "") {
            description.style.border = '2px solid red';
            create.style.backgroundColor = '#e4e4e4';
            create.style.color = '#b2abab';
            create.style.cursor = 'default';
            create.removeAttribute('type');
            create.setAttribute('type', 'button');
        } else{
            description.style.border = '2px solid rgba(0,0,0,0)';
            description.classList.remove('pin');
            if(projectName.value != "" && startDate.value != "" && endDate.value != "" && description.value != "" && uploadImg.value != "") {
                create.style.backgroundColor = '#000';
                create.style.color = '#eaeaea';
                create.style.cursor = 'pointer';
                create.removeAttribute('type');
                create.setAttribute('type', 'submit');
            }
        }
    })
})


uploadImg.addEventListener('click', function() {
    
    uploadImg.classList.add('pin');
    uploadImg.removeAttribute('placeholder');
    if(startDate.classList.contains('pin') && startDate.value == "") {
        startDate.style.border = '2px solid red';
    }
    if(endDate.classList.contains('pin') && endDate.value == "") {
        endDate.style.border = '2px solid red';
    }
    if(description.classList.contains('pin') && description.value == "") {
        description.style.border = '2px solid red';
    }
    if(projectName.classList.contains('pin') && projectName.value == "") {
        projectName.style.border = '2px solid red';
    }
    uploadImg.addEventListener('input', function() {
        if(uploadImg.value == "") {
            document.querySelector('.boxField').style.border = '2px solid red';
            document.querySelector('.choose').style.backgroundColor = '#e4e4e4';
            document.querySelector('.choose').style.color = '#b2abab';
            document.querySelector('.choose').innerHTML = 'choose';
            document.querySelector('.linkLogo').setAttribute("src", "../public/img/link.png");
            create.style.backgroundColor = '#e4e4e4';
            create.style.color = '#b2abab';
            create.style.cursor = 'default';
            create.removeAttribute('type');
            create.setAttribute('type', 'button');
        } else{
            document.querySelector('.boxField').style.border = '2px solid rgba(0,0,0,0)';
            document.querySelector('.choose').style.backgroundColor = 'black';
            document.querySelector('.choose').style.color = '#eaeaea';
            document.querySelector('.choose').innerHTML = 'choosed';
            document.querySelector('.linkLogo').setAttribute("src", "../public/img/check.png");
            document.querySelector('.linkLogo').style.filter = 'drop-shadow(0 0 3px rgba(120, 120, 120, .7))';
            uploadImg.classList.remove('pin');
            if(projectName.value != "" && startDate.value != "" && endDate.value != "" && description.value != "" && uploadImg.value != "") {
                create.style.backgroundColor = '#000';
                create.style.color = '#eaeaea';
                create.style.cursor = 'pointer';
                create.removeAttribute('type');
                create.setAttribute('type', 'submit');
            }
        }
    })
})


create.addEventListener('click', function() {

    if(projectName.value != "" && startDate.value != "" && endDate.value != "" && description.value != "" && uploadImg.value != "") {
        
    } else {
        alert("Semua field wajib diisi!");
    }
})

