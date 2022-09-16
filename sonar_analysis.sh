containerIP=`docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' sonarqube`
echo $containerIP

docker run --rm \
  -v "$(pwd):/usr/src" \
  sonarsource/sonar-scanner-cli \
 -Dproject.settings=./sonar-project.properties \
 -Dsonar.host.url=http://$containerIP:9000 \
 -Dsonar.login=