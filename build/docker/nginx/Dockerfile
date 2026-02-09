FROM nginx

COPY build/docker/nginx/nginx.conf /etc/nginx/nginx.conf
COPY build/docker/nginx/templates/ /etc/nginx/templates/
COPY web/crm/dist/ /usr/share/nginx/html/
COPY web/project-page/dist/ /usr/share/nginx/html/project-page/