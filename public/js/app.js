function confirmPasswordLength() {
  var passw = $('#password');
  var notify = $('<small style="color:red; display:block;">');
  passw.on('input', function (e) {
    if (passw.val().length > 7) {
      notify.text('');
      $('button:submit').removeAttr('disabled');
    } else {
      notify.text('Password must be at least 8 characters');
      $('button:submit').attr('disabled', 'disabled');
      passw.parent().append(notify);
    }
  });
}

function confirmPasswordsMatch() {
  //confirm password matches (reg page only)
  var passw = $('#password');
  var notify = $('<small style="color:red; display:block;">');
  var confirmPw = $('#confirm-password');
  confirmPw.on('input', function (e) {
    if (passw.val() === e.target.value) {
      notify.text('');
      $('button:submit').removeAttr('disabled');
    } else {
      notify.text('Passwords do not match');
      $('button:submit').attr('disabled', 'disabled');
      confirmPw.parent().append(notify);
    }
  });
}