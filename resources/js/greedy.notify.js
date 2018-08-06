/*
CreateNotify(
	'<span class="fa fa-info"></span>', //icon
	'Benachrichtigung', //title
	'Ich bin eine Benachrichtigung', //content
	'info', //type
	'#main_notifies', //parent
	'true' //autoclose true or false (optional parameter)
);
*/

function CreateNotify(icon, title, content, type, parent, autoclose) {
    var notify_main = $("<div/>").addClass("js-notify " + type).appendTo(parent);
    $(notify_main).css('display', 'none');
    var notify_icon = $("<div/>").addClass("icon-content").appendTo(notify_main);
    var notify_content = $("<div/>").addClass("content").appendTo(notify_main);
    var notify_content_title = $("<div/>").addClass("title").appendTo(notify_content);
    var notify_content_subtitle = $("<div/>").addClass("subtitle").appendTo(notify_content);

    if (typeof autoclose === 'undefined' || autoclose === null) {
        autoclose = 'true';
    }

    $(notify_icon).html(icon);
    $(notify_content_title).html(title);
    $(notify_content_subtitle).html(content);

    if ($(parent + '-msmall').length) {
        var notify_main_small = $("<div/>").addClass("js-notify-small " + type).appendTo(parent + '-msmall');
        $(notify_main_small).css('display', 'none');
        var notify_icon_small = $("<div/>").addClass("icon-content").appendTo(notify_main_small);

        $(notify_icon_small).html(icon);

        $(notify_main_small).show('fade', 'slow', function () {
            if (autoclose == 'true') {
                setTimeout(function () {
                    $(notify_main_small).hide('fade', 'slow');
                }, 5000);
            }
        });
    }

    $(notify_main).show('fade', 'slow', function () {
        if (autoclose == 'true') {
            setTimeout(function () {
                $(notify_main).hide('fade', 'slow');
            }, 5000);
        }
    });
}

function CreateNotifyWithID(id, icon, title, content, type, parent, autoclose) {
    var notify_main = $("<div/>").addClass("js-notify " + type).appendTo(parent);
    $(notify_main).attr('id', id);
    $(notify_main).css('display', 'none');
    var notify_icon = $("<div/>").addClass("icon-content").appendTo(notify_main);
    var notify_content = $("<div/>").addClass("content").appendTo(notify_main);
    var notify_content_title = $("<div/>").addClass("title").appendTo(notify_content);
    var notify_content_subtitle = $("<div/>").addClass("subtitle").appendTo(notify_content);

    if (typeof autoclose === 'undefined' || autoclose === null) {
        autoclose = 'true';
    }

    $(notify_icon).html(icon);
    $(notify_content_title).html(title);
    $(notify_content_subtitle).html(content);

    $(notify_main).show('fade', 'slow', function () {
        if (autoclose == 'true') {
            setTimeout(function () {
                $(notify_main).hide('fade', 'slow', function () {
                    $(notify_main).remove();
                });
            }, 5000);
        }
    });
}