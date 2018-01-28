//regular app features
$(function () {
	//like functionality
	var actions = { like: 'Like', unlike: 'Unlike' };
	var likeBtn = $('#like-btn');
	var like = $('#likes');
	var likeCount = parseInt(like.text());
	likeBtn.on('click', function (e) {
		var $this = $(this);
		var currentAction = $this.data('action').toLowerCase();
		var nextAction = null;

		e.preventDefault();
		if (currentAction === actions.unlike.toLowerCase()) {
			likeCount--;
			nextAction = actions.like.toLowerCase();
		} else {
			likeCount++;
			nextAction = actions.unlike.toLowerCase();
		}
		//set next action
		$this.data('action', nextAction);
		var update = {
			action: currentAction,
		};
		$.post(
			'/article/' + $this.data('id') +'/like',
			JSON.stringify(update),
			function (data) {
				like.text(data.likes);
				$this.find('i').toggleClass("fa-thumbs-up").toggleClass("fa-thumbs-down");
				likeBtn.toggleClass('btn-secondary').toggleClass('btn-success');
			}
		);
	});
});

//registration and login
$(function () {
	//check password length (reg and login pages)
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

	//confirm password matches (reg page only)
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
});