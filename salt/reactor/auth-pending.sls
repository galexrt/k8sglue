{# Inspired by https://www.reddit.com/r/saltstack/comments/3jeggp/automating_key_acceptance/cuorkqc #}
{% if 'act' in data and data['act'] == 'pend' %}
minion_add:
{# TODO use safe_accept sate #}
  wheel.key.accept:
 - match: {{ data['id'] }}
{% endif %}
