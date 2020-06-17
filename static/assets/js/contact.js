$(function() {

        $('#toggle').on('click', function() {

            if($("#slide").hasClass('visible')){
                $("#slide").removeAttr( 'style' );
                $("#slide").removeClass('visible');
                $("#toggle").css('background-image','url(../../images/arrow-leaf-right.png)');
            }
            else {
                $("#slide").css('left', '0' );
                $("#slide").addClass('visible');
                $("#toggle").css('background-image','url(../../images/arrow-leaf-left.png)');
            }
        });

        //If an email address has been entered successfully reset the Slider
        $('#subscribe-success').on('click', function() {
            $("#slide").removeAttr( 'style' );
            $("#subscribe-email").val('');
            $("#fm-subscribe-email").show();
            $("#subscribe-text").show();
            $("#subscribe-success").hide();
        });

        $('#subscribe-email-btn').on('click', function() {
            //hide error
            $("#invalid-email").hide("fast");
            $("#invalid-name").hide("fast");

            if(!isValidName($("#subscribe-name").val()))
            {
                //show error
                $("#invalid-name").show("slow");
                return;
            }

            if(!isValidEmail($("#subscribe-email").val()))
            {
                //show error
                $("#invalid-email").show("slow");
                return;
            }

            // Get email address values from elements on the page:
            var formElements = $("#fm-subscribe-email").serializeArray();

            var subscribeForm = {};
            for (var i = 0; i < formElements.length; i++) {
                subscribeForm[formElements[i].name] = formElements[i].value;
            }

            var dataJSON = JSON.stringify(subscribeForm);

            $.ajax({
                type: "POST",
                url: "subscribe-email",
                data: dataJSON,
                success: function(){},
                dataType: "json",
                contentType : "application/json"
            });
            $("#subscribe-success").show("slow");
            $("#slide").css('left','0');
            $("#fm-subscribe-email").hide();
            $("#subscribe-text").hide();


        });

        $('#submit-email-btn').on('click', function() {

        // Get some values from elements on the page:
        var formElements = $("#fm-submit-email").serializeArray();

        var contactForm = {};
        for (var i = 0; i < formElements.length; i++) {
            contactForm[formElements[i].name] = formElements[i].value;
        }

        var dataJSON = JSON.stringify(contactForm);

        $.ajax({
            type: "POST",
            url: "send-email",
            data: dataJSON,
            success: function(){},
            dataType: "json",
            contentType : "application/json"
        });

        //reset form
        $("#fm-submit-email").find('select')
            .val($('option:first').val());

        $("#fm-submit-email").find('input,textarea')
            .each(function() {

                var i = $(this),
                    x;

                i.removeClass('polyfill-placeholder');

                switch (this.type) {

                    case 'button':
                    case 'reset':
                        break;

                    case 'checkbox':
                    case 'radio':
                        i.attr('checked', i.attr('defaultValue'));
                        break;

                    case 'text':
                    case 'textarea':
                        i.val(i.attr('defaultValue'));

                        if (i.val() == '') {
                            i.addClass('polyfill-placeholder');
                            i.val(i.attr('placeholder'));
                        }

                        break;

                    default:
                        i.val(i.attr('defaultValue'));
                        break;

                }
            });

        //hide input fields
        $("#fm-submit-email").hide("slow");

        //Show notification
        $("#notification").show();

    });
});

function isValidEmail(email)
{
    if (/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,10})+$/.test(email))
    {
        return (true)
    }
    return (false)
}

function isValidName(email)
{
    if (/^(([A-Za-z]+[\-\']?)*([A-Za-z]+)?\s)*([A-Za-z]+[\-\']?)*([A-Za-z]+)?$/.test(email))
    {
        return (true)
    }
    return (false)
}

